/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package api

// CallbackNotification define blockchain callback notification
type CallbackNotification struct {
	Blockchain string `json:"blockchain"`
	ID         string `json:"id"`
	BlockNum   int64  `json:"block_num"`
	TxID       string `json:"tx_id"`
	Time       int64  `json:"time"`
}

// CallbackResponse define blockchain callback response
type CallbackResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}
