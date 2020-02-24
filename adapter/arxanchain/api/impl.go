/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package api

import (
	"fmt"

	"github.com/csiabb/blockchain-adapter/adapter/arxanchain/structs"
)

// CommitTransaction implement blockchain adapter interface
func (c *ArxanchainClient) CommitTransaction(commit interface{}) (interface{}, error) {
	if nil == commit {
		return nil, fmt.Errorf("Commit blockchain transaction data is nil")
	}

	r, ok := commit.(structs.CommonRequest)
	if !ok {
		return nil, fmt.Errorf("request type is error")
	}

	if nil == r.Request {
		return nil, fmt.Errorf("request data is nil")
	}

	switch r.Action {
	case structs.CreateAccount:
		createReqeust, ok := r.Request.(structs.CreateAccountRequest)
		if !ok {
			return nil, fmt.Errorf("create account request data type is error")
		}
		return c.CreateAccount(&createReqeust)
	}

	return nil, fmt.Errorf("Action %s is unsupport", r.Action)
}

// QueryTransaction implement blockchain adapter interface
func (c *ArxanchainClient) QueryTransaction(query interface{}) (interface{}, error) {
	if nil == query {
		return nil, fmt.Errorf("Query blockchain transaction request is nil")
	}

	r, ok := query.(structs.CommonRequest)
	if !ok {
		return nil, fmt.Errorf("request type is error")
	}

	if nil == r.Request {
		return nil, fmt.Errorf("request data is nil")
	}

	switch r.Action {
	case structs.CreateAsset:
		createReqeust, ok := r.Request.(structs.AssetRegisterRequest)
		if !ok {
			return nil, fmt.Errorf("create asset request data type is error")
		}
		return c.RegisteAsset(&createReqeust)
	case structs.QueryAsset:
		queryRequest, ok := r.Request.(structs.QueryAssetDetailRequest)
		if !ok {
			return nil, fmt.Errorf("query asset request data type is error")
		}
		return c.QueryAssetDetail(&queryRequest)
	}

	return nil, nil
}
