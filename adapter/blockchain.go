/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package adapter

//go:generate mockgen -destination=adapter/mock/mock_adapter.go -package=mock_adapter github.com/csiabb/blockchain-adapter/adapter BlockchainAdapter

// BlockchainAdapter adapter of blockchain
type BlockchainAdapter interface {
	CreateAccount(*CreateAccountReq) (*BlockchainResponse, error)
	PublicityData(*PublicityDataReq) (*BlockchainResponse, error)
}

// CreateAccountReq request of create account
type CreateAccountReq struct {
	AccountID string
}

// PublicityDataReq request of publicity data
type PublicityDataReq struct {
	AccountID string
	Publicity string
}

// BlockchainResponse response of blockchain
type BlockchainResponse struct {
	ID string `json:"id"`
}
