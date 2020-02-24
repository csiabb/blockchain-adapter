/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package api

import (
	"testing"

	rhttp "github.com/csiabb/blockchain-adapter/adapter/arxanchain/rest/http"
)

func TestCommitTxResErr(t *testing.T) {
	cfg := &rhttp.Config{
		Host:      "https://aboxtest.arxanchain.com",
		APIKey:    "XxEVWIAD1579163844",
		APISecret: "nXAVIKEMHUONGXREDDBAAUDXIPWQIGQZ",
	}

	c, err := NewArxanchainClient(cfg)
	if nil != err {
		t.Error("test new arxanchain client error: ", err)
	}

	res := "test"
	_, err = c.CommitTransaction(res)

	if nil == err {
		t.Error("request error: ", err)
	}
}
