/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/csiabb/blockchain-adapter/adapter/arxanchain/common"
	rhttp "github.com/csiabb/blockchain-adapter/adapter/arxanchain/rest/http"
	"github.com/csiabb/blockchain-adapter/adapter/arxanchain/structs"
	"github.com/csiabb/blockchain-adapter/common/log"
)

var (
	logger = log.MustGetLogger("arxanchain")
)

// ArxanchainClient aranchain block client
type ArxanchainClient struct {
	c *rhttp.Client
}

// NewArxanchainClient new arxanchain client
func NewArxanchainClient(cfg *rhttp.Config) (*ArxanchainClient, error) {
	if nil == cfg {
		return nil, fmt.Errorf("config is nil")
	}

	if !cfg.Enabled {
		return nil, fmt.Errorf("arxanchain is disabled")
	}

	if "" == cfg.Endpoint {
		return nil, fmt.Errorf("Missing endpoint config")
	}

	if "" == cfg.APIKey {
		return nil, fmt.Errorf("Missing api key config")
	}

	if "" == cfg.APISecret {
		return nil, fmt.Errorf("Missing api secret config")
	}

	c, err := rhttp.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return &ArxanchainClient{c: c}, nil
}

func (ac *ArxanchainClient) addSignatureHeader(header *http.Header, path, method string) error {
	if nil == header {
		return fmt.Errorf("header is nil")
	}

	header.Add(structs.APIKeyHeader, ac.c.Cfg.APIKey)

	sigData := &common.SignatureData{
		Secret:        ac.c.Cfg.APISecret,
		SignAlgo:      structs.HMACSHA256,
		RequestPath:   path,
		RequestMethod: method,
		Timestamp:     time.Now().Unix(),
	}

	signature, err := common.Signature(sigData)
	if nil != err {
		return err
	}

	header.Add(structs.SignatureTimestamp, fmt.Sprintf("%d", sigData.Timestamp))
	header.Add(structs.SignatureMethod, sigData.SignAlgo)
	header.Add(structs.SignatureInfo, signature)

	return nil
}
