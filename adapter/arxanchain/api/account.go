/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/csiabb/blockchain-adapter/adapter/arxanchain/structs"
)

// CreateAccount ...
func (c *ArxanchainClient) CreateAccount(body *structs.CreateAccountRequest) (result *structs.AccountResponse, err error) {
	if body == nil {
		err = fmt.Errorf("request payload is null")
		return
	}

	header := http.Header{}
	err = c.addSignatureHeader(&header, "", structs.PostMethod)
	if nil != err {
		return
	}

	// Build http request
	r := c.NewRequest(structs.PostMethod, structs.CreateAccountURL)
	r.SetHeaders(header)
	r.SetBody(body)

	// Do http request
	_, resp, err := RequireOK(c.DoRequest(r))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Parse http response
	var respBody structs.CommonResponse
	if err = DecodeBody(resp, &respBody); err != nil {
		return
	}

	if respBody.Code != structs.SuccCode {
		err = fmt.Errorf(respBody.Message)
		return
	}

	payloadBytes, err := json.Marshal(respBody.Payload)
	if nil != err {
		return
	}

	err = json.Unmarshal(payloadBytes, &result)
	return
}
