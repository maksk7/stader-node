/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.4.0]

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
// Code generated by fastssz. DO NOT EDIT.
// Hash: 1a4a8518e54915bb492ddfc9d58736ddc7a3f6a375a430f370a808b609511640
// Version: 0.1.2
package eth2

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the DepositDataNoSignature object
func (d *DepositDataNoSignature) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(d)
}

// MarshalSSZTo ssz marshals the DepositDataNoSignature object to a target array
func (d *DepositDataNoSignature) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'PublicKey'
	if size := len(d.PublicKey); size != 48 {
		err = ssz.ErrBytesLengthFn("DepositDataNoSignature.PublicKey", size, 48)
		return
	}
	dst = append(dst, d.PublicKey...)

	// Field (1) 'WithdrawalCredentials'
	if size := len(d.WithdrawalCredentials); size != 32 {
		err = ssz.ErrBytesLengthFn("DepositDataNoSignature.WithdrawalCredentials", size, 32)
		return
	}
	dst = append(dst, d.WithdrawalCredentials...)

	// Field (2) 'Amount'
	dst = ssz.MarshalUint64(dst, d.Amount)

	return
}

// UnmarshalSSZ ssz unmarshals the DepositDataNoSignature object
func (d *DepositDataNoSignature) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 88 {
		return ssz.ErrSize
	}

	// Field (0) 'PublicKey'
	if cap(d.PublicKey) == 0 {
		d.PublicKey = make([]byte, 0, len(buf[0:48]))
	}
	d.PublicKey = append(d.PublicKey, buf[0:48]...)

	// Field (1) 'WithdrawalCredentials'
	if cap(d.WithdrawalCredentials) == 0 {
		d.WithdrawalCredentials = make([]byte, 0, len(buf[48:80]))
	}
	d.WithdrawalCredentials = append(d.WithdrawalCredentials, buf[48:80]...)

	// Field (2) 'Amount'
	d.Amount = ssz.UnmarshallUint64(buf[80:88])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the DepositDataNoSignature object
func (d *DepositDataNoSignature) SizeSSZ() (size int) {
	size = 88
	return
}

// HashTreeRoot ssz hashes the DepositDataNoSignature object
func (d *DepositDataNoSignature) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(d)
}

// HashTreeRootWith ssz hashes the DepositDataNoSignature object with a hasher
func (d *DepositDataNoSignature) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'PublicKey'
	if size := len(d.PublicKey); size != 48 {
		err = ssz.ErrBytesLengthFn("DepositDataNoSignature.PublicKey", size, 48)
		return
	}
	hh.PutBytes(d.PublicKey)

	// Field (1) 'WithdrawalCredentials'
	if size := len(d.WithdrawalCredentials); size != 32 {
		err = ssz.ErrBytesLengthFn("DepositDataNoSignature.WithdrawalCredentials", size, 32)
		return
	}
	hh.PutBytes(d.WithdrawalCredentials)

	// Field (2) 'Amount'
	hh.PutUint64(d.Amount)

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the DepositDataNoSignature object
func (d *DepositDataNoSignature) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(d)
}

// MarshalSSZ ssz marshals the DepositData object
func (d *DepositData) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(d)
}

// MarshalSSZTo ssz marshals the DepositData object to a target array
func (d *DepositData) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'PublicKey'
	if size := len(d.PublicKey); size != 48 {
		err = ssz.ErrBytesLengthFn("DepositData.PublicKey", size, 48)
		return
	}
	dst = append(dst, d.PublicKey...)

	// Field (1) 'WithdrawalCredentials'
	if size := len(d.WithdrawalCredentials); size != 32 {
		err = ssz.ErrBytesLengthFn("DepositData.WithdrawalCredentials", size, 32)
		return
	}
	dst = append(dst, d.WithdrawalCredentials...)

	// Field (2) 'Amount'
	dst = ssz.MarshalUint64(dst, d.Amount)

	// Field (3) 'Signature'
	if size := len(d.Signature); size != 96 {
		err = ssz.ErrBytesLengthFn("DepositData.Signature", size, 96)
		return
	}
	dst = append(dst, d.Signature...)

	return
}

// UnmarshalSSZ ssz unmarshals the DepositData object
func (d *DepositData) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 184 {
		return ssz.ErrSize
	}

	// Field (0) 'PublicKey'
	if cap(d.PublicKey) == 0 {
		d.PublicKey = make([]byte, 0, len(buf[0:48]))
	}
	d.PublicKey = append(d.PublicKey, buf[0:48]...)

	// Field (1) 'WithdrawalCredentials'
	if cap(d.WithdrawalCredentials) == 0 {
		d.WithdrawalCredentials = make([]byte, 0, len(buf[48:80]))
	}
	d.WithdrawalCredentials = append(d.WithdrawalCredentials, buf[48:80]...)

	// Field (2) 'Amount'
	d.Amount = ssz.UnmarshallUint64(buf[80:88])

	// Field (3) 'Signature'
	if cap(d.Signature) == 0 {
		d.Signature = make([]byte, 0, len(buf[88:184]))
	}
	d.Signature = append(d.Signature, buf[88:184]...)

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the DepositData object
func (d *DepositData) SizeSSZ() (size int) {
	size = 184
	return
}

// HashTreeRoot ssz hashes the DepositData object
func (d *DepositData) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(d)
}

// HashTreeRootWith ssz hashes the DepositData object with a hasher
func (d *DepositData) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'PublicKey'
	if size := len(d.PublicKey); size != 48 {
		err = ssz.ErrBytesLengthFn("DepositData.PublicKey", size, 48)
		return
	}
	hh.PutBytes(d.PublicKey)

	// Field (1) 'WithdrawalCredentials'
	if size := len(d.WithdrawalCredentials); size != 32 {
		err = ssz.ErrBytesLengthFn("DepositData.WithdrawalCredentials", size, 32)
		return
	}
	hh.PutBytes(d.WithdrawalCredentials)

	// Field (2) 'Amount'
	hh.PutUint64(d.Amount)

	// Field (3) 'Signature'
	if size := len(d.Signature); size != 96 {
		err = ssz.ErrBytesLengthFn("DepositData.Signature", size, 96)
		return
	}
	hh.PutBytes(d.Signature)

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the DepositData object
func (d *DepositData) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(d)
}

// MarshalSSZ ssz marshals the SigningRoot object
func (s *SigningRoot) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SigningRoot object to a target array
func (s *SigningRoot) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'ObjectRoot'
	if size := len(s.ObjectRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("SigningRoot.ObjectRoot", size, 32)
		return
	}
	dst = append(dst, s.ObjectRoot...)

	// Field (1) 'Domain'
	if size := len(s.Domain); size != 32 {
		err = ssz.ErrBytesLengthFn("SigningRoot.Domain", size, 32)
		return
	}
	dst = append(dst, s.Domain...)

	return
}

// UnmarshalSSZ ssz unmarshals the SigningRoot object
func (s *SigningRoot) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 64 {
		return ssz.ErrSize
	}

	// Field (0) 'ObjectRoot'
	if cap(s.ObjectRoot) == 0 {
		s.ObjectRoot = make([]byte, 0, len(buf[0:32]))
	}
	s.ObjectRoot = append(s.ObjectRoot, buf[0:32]...)

	// Field (1) 'Domain'
	if cap(s.Domain) == 0 {
		s.Domain = make([]byte, 0, len(buf[32:64]))
	}
	s.Domain = append(s.Domain, buf[32:64]...)

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SigningRoot object
func (s *SigningRoot) SizeSSZ() (size int) {
	size = 64
	return
}

// HashTreeRoot ssz hashes the SigningRoot object
func (s *SigningRoot) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SigningRoot object with a hasher
func (s *SigningRoot) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'ObjectRoot'
	if size := len(s.ObjectRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("SigningRoot.ObjectRoot", size, 32)
		return
	}
	hh.PutBytes(s.ObjectRoot)

	// Field (1) 'Domain'
	if size := len(s.Domain); size != 32 {
		err = ssz.ErrBytesLengthFn("SigningRoot.Domain", size, 32)
		return
	}
	hh.PutBytes(s.Domain)

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the SigningRoot object
func (s *SigningRoot) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}

// MarshalSSZ ssz marshals the VoluntaryExit object
func (v *VoluntaryExit) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(v)
}

// MarshalSSZTo ssz marshals the VoluntaryExit object to a target array
func (v *VoluntaryExit) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Epoch'
	dst = ssz.MarshalUint64(dst, v.Epoch)

	// Field (1) 'ValidatorIndex'
	dst = ssz.MarshalUint64(dst, v.ValidatorIndex)

	return
}

// UnmarshalSSZ ssz unmarshals the VoluntaryExit object
func (v *VoluntaryExit) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 16 {
		return ssz.ErrSize
	}

	// Field (0) 'Epoch'
	v.Epoch = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'ValidatorIndex'
	v.ValidatorIndex = ssz.UnmarshallUint64(buf[8:16])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the VoluntaryExit object
func (v *VoluntaryExit) SizeSSZ() (size int) {
	size = 16
	return
}

// HashTreeRoot ssz hashes the VoluntaryExit object
func (v *VoluntaryExit) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(v)
}

// HashTreeRootWith ssz hashes the VoluntaryExit object with a hasher
func (v *VoluntaryExit) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Epoch'
	hh.PutUint64(v.Epoch)

	// Field (1) 'ValidatorIndex'
	hh.PutUint64(v.ValidatorIndex)

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the VoluntaryExit object
func (v *VoluntaryExit) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(v)
}
