package stader_backend

type PreSignCheckApiRequestType struct {
	ValidatorPublicKey string `json:"validatorPublicKey"`
}

type PreSignCheckApiResponseType struct {
	Value bool `json:"value"`
}

type PreSignSendApiResponseType struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type PreSignSendApiRequestType struct {
	Message struct {
		Epoch          uint64 `json:"epoch"`
		ValidatorIndex uint64 `json:"validator_index"`
	} `json:"message"`
	MessageHash        []byte `json:"messageHash"`
	Signature          []byte `json:"signature"`
	ValidatorPublicKey string `json:"validatorPublicKey"`
}

type PublicKeyApiResponse struct {
	Value string `json:"value"`
}
