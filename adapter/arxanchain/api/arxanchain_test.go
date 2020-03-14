/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package api

import (
	"testing"

	rhttp "github.com/csiabb/blockchain-adapter/adapter/arxanchain/rest/http"
)

func TestNewArxanchainClient(t *testing.T) {
	cfg := &rhttp.Config{
		Enabled:   true,
		Endpoint:  "https://aboxtest.arxanchain.com",
		APIKey:    "XxEVWIAD1579163843",
		APISecret: "nXAVIKEMHUONGXREDDBAAUDXIPWQIGQZ",
	}

	_, err := NewArxanchainClient(cfg)
	if nil != err {
		t.Error("test new arxanchain client error: ", err)
	}
}
