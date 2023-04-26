/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [0.3.0-beta]

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package guardian

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/services/state"
	"github.com/stader-labs/stader-node/stader/guardian/collector"

	"github.com/fatih/color"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/utils/log"
)

// Config
var tasksInterval, _ = time.ParseDuration("5s")
var taskCooldown, _ = time.ParseDuration("10s")

const (
	MaxConcurrentEth1Requests = 200

	ErrorColor   = color.FgRed
	UpdateColor  = color.FgBlue
	MetricsColor = color.FgHiYellow
)

// Register guardian command
func RegisterCommands(app *cli.App, name string, aliases []string) {
	app.Commands = append(app.Commands, cli.Command{
		Name:    name,
		Aliases: aliases,
		Usage:   "Run Stader guardian activity daemon",
		Action: func(c *cli.Context) error {
			return run(c)
		},
	})
}

// Run daemon
func run(c *cli.Context) error {

	// Initialize error logger
	errorLog := log.NewColorLogger(ErrorColor)
	updateLog := log.NewColorLogger(UpdateColor)

	// Configure
	configureHTTP()

	cfg, err := services.GetConfig(c)
	if err != nil {
		return err
	}
	ec, err := services.GetEthClient(c)
	if err != nil {
		return err
	}
	bc, err := services.GetBeaconClient(c)
	if err != nil {
		return err
	}

	stateCache := collector.NewStateCache()
	w, err := services.GetWallet(c)
	if err != nil {
		return err
	}

	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return err
	}
	// Run metrics loop
	go func() {
		m, err := state.NewNetworkStateManager(cfg, ec, bc, &updateLog)
		if err != nil {
			panic(err)
		}

		for {
			// Check the EC status
			err := services.WaitEthClientSynced(c, false) // Force refresh the primary / fallback EC status
			if err != nil {
				errorLog.Println("WaitEthClientSynced ", err)
				time.Sleep(taskCooldown)
				continue
			}

			// Check the BC status
			err = services.WaitBeaconClientSynced(c, false) // Force refresh the primary / fallback BC status
			if err != nil {
				errorLog.Println("WaitBeaconClientSynced ", err)
				time.Sleep(taskCooldown)
				continue
			}

			networkStateCache, err := updateNetworkStateCache(m, nodeAccount.Address)
			if err != nil {
				errorLog.Println("updateNetworkStateCache ", err)
				time.Sleep(taskCooldown)
				continue
			}
			stateCache.UpdateState(networkStateCache)
			time.Sleep(tasksInterval)
		}
	}()

	err = runMetricsServer(c, log.NewColorLogger(MetricsColor), stateCache)
	if err != nil {
		errorLog.Println(err)
	}
	return nil
}

// Configure HTTP transport settings
func configureHTTP() {

	// The guardian daemon makes a large number of concurrent RPC requests to the Eth1 client
	// The HTTP transport is set to cache connections for future re-use equal to the maximum expected number of concurrent requests
	// This prevents issues related to memory consumption and address allowance from repeatedly opening and closing connections
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = MaxConcurrentEth1Requests
}

func updateNetworkStateCache(m *state.NetworkStateManager, nodeAddress common.Address) (*state.NetworkStateCache, error) {
	// Get the networkStateCache of the network
	networkStateCache, err := m.GetHeadStateForNode(nodeAddress)
	if err != nil {
		return nil, fmt.Errorf("error updating network state: %w", err)
	}
	return networkStateCache, nil
}
