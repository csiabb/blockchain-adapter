/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package service

import (
	"github.com/csiabb/blockchain-adapter/common/metadata"
	"github.com/csiabb/blockchain-adapter/config"
)

// Server interface ...
type Server interface {
	Start() (err error)
	Shutdown()
}

// NewUninitializedServer creates a Server instance without initializing it
func NewUninitializedServer(c *config.SrvcCfg, sVer *metadata.Version) (Server, error) {
	if BlockchainAdapter == nil {
		BlockchainAdapter = &ServerImpl{
			config:  c,
			version: sVer,
		}
	}
	return BlockchainAdapter, nil
}

// NewServer create blockchain adapter server
func NewServer(c *config.SrvcCfg, sVer *metadata.Version) (Server, error) {
	if BlockchainAdapter == nil {
		BlockchainAdapter = &ServerImpl{
			config:  c,
			version: sVer,
		}
		err := BlockchainAdapter.init()
		if err != nil {
			logger.Errorf("Failed to initialize unite did server, %+v", err)
			return nil, err
		}
	}
	return BlockchainAdapter, nil
}
