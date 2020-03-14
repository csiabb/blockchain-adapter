/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/csiabb/blockchain-adapter/adapter"
	"github.com/csiabb/blockchain-adapter/adapter/arxanchain/rest"
	rhttp "github.com/csiabb/blockchain-adapter/adapter/arxanchain/rest/http"
	"github.com/csiabb/blockchain-adapter/adapter/arxanchain/structs"
)

// CreateAccount ...
func (ac *ArxanchainClient) CreateAccount(param *adapter.CreateAccountReq) (result *adapter.BlockchainResponse, err error) {
	if nil == param {
		err = fmt.Errorf("param is nil")
		return
	}

	body := &structs.CreateAccountRequest{UID: param.AccountID}

	header := http.Header{}
	err = ac.addSignatureHeader(&header, structs.CreateAccountURL, http.MethodPost)
	if nil != err {
		return
	}
	header.Add("Content-Type", "application/json")

	// Build http request
	r := ac.c.NewRequest(http.MethodPost, structs.CreateAccountURL)
	r.SetHeaders(header)
	r.SetBody(body)

	// Do http request
	_, resp, err := rhttp.RequireOK(ac.c.DoRequest(r))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Parse http response
	var respBody structs.ArxanResponse
	if err = rhttp.DecodeBody(resp, &respBody); err != nil {
		return
	}

	if respBody.Code != rest.SuccCode {
		err = fmt.Errorf("Code: %d, Message: %s", respBody.Code, respBody.Message)
		return
	}

	payloadBytes, err := json.Marshal(respBody.Payload)
	if nil != err {
		return
	}

	var respAccount *structs.AccountResponse
	err = json.Unmarshal(payloadBytes, &respAccount)
	if nil != err {
		return
	}

	result = &adapter.BlockchainResponse{ID: respAccount.DID}
	return
}
