/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/csiabb/blockchain-adapter/adapter/arxanchain/structs"
)

// CommitTransaction implement blockchain adapter interface
func (c *ArxanchainClient) CommitTransaction(commit interface{}) (interface{}, error) {
	if nil == commit {
		return nil, fmt.Errorf("Commit blockchain transaction data is nil")
	}
	return nil, nil
}

// QueryTransaction implement blockchain adapter interface
func (c *ArxanchainClient) QueryTransaction(query interface{}) (interface{}, error) {
	if nil == query {
		return nil, fmt.Errorf("Query blockchain transaction request is nil")
	}
	return nil, nil
}

func (c *ArxanchainClient) addSignatureHeader(header *http.Header, path, method string) error {
	if nil == header {
		return fmt.Errorf("header is nil")
	}
	header.Add(structs.APIKeyHeader, c.config.APIKey)

	sigData := &SignatureData{
		Secret:        c.config.APISecret,
		SignAlgo:      structs.HMACSHA256,
		RequestPath:   path,
		RequestMethod: method,
		Timestamp:     time.Now().Unix(),
	}

	signature, err := Signature(sigData)
	if nil != err {
		return err
	}

	header.Add(SignatureTimestamp, fmt.Sprintf("%d", sigData.Timestamp))
	header.Add(SignatureMethod, sigData.SignAlgo)
	header.Add(SignatureInfo, signature)

	return nil
}
