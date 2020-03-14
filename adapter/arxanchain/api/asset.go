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

// PublicityData ...
func (ac *ArxanchainClient) PublicityData(param *adapter.PublicityDataReq) (result *adapter.BlockchainResponse, err error) {
	if nil == param {
		err = fmt.Errorf("param is nil")
		return
	}
	body := &structs.AssetRegisterRequest{
		OwnerDID:   param.AccountID,
		AssetName:  "publicity-data",
		Visibility: "private",
		Intro:      "publicity-data of " + param.AccountID,
		Metadata:   param.Publicity,
	}

	header := http.Header{}
	err = ac.addSignatureHeader(&header, structs.CreateAssetURL, http.MethodPost)
	if nil != err {
		return
	}
	header.Add("Content-Type", "application/json")

	// Build http request
	r := ac.c.NewRequest(http.MethodPost, structs.CreateAssetURL)
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

	var respAssest *structs.AssetRegisterResponse
	err = json.Unmarshal(payloadBytes, &respAssest)
	if nil != err {
		return
	}

	result = &adapter.BlockchainResponse{ID: respAssest.AssetDID}
	return
}
