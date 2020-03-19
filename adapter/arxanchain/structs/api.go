/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package structs

// ArxanResponse common rest response
type ArxanResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}

// CreateAccountRequest ...
type CreateAccountRequest struct {
	Phone       string   `json:"phone"`
	Email       string   `json:"email"`
	UID         string   `json:"uid"`
	Tags        []string `json:"tags"`
	CallbackURL string   `json:"callback_url"`
}

// AccountResponse ...
type AccountResponse struct {
	DID  string   `json:"did"`
	Tags []string `json:"tags"`
}

// AssetRegisterRequest ...
type AssetRegisterRequest struct {
	OwnerDID    string   `json:"owner_did"`
	AssetName   string   `json:"asset_name"`
	Intro       string   `json:"intro"`
	Visibility  string   `json:"visibility"`
	Category    string   `json:"category"`
	Source      string   `json:"source"`
	Tags        []string `json:"tags"`
	Metadata    string   `json:"metadata"`
	CallbackURL string   `json:"callback_url"`
}

// AssetRegisterResponse ...
type AssetRegisterResponse struct {
	AssetDID string `json:"asset_did"`
}

// CallBackRequest ...
type CallBackRequest struct {
	ID          string `json:"id" binding:"required"`
	TxID        string `json:"tx_id" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Status      string `json:"status" binding:"required"`
	BlockNumber uint64 `json:"block_number" binding:"required"`
	OnChainTime int64  `json:"on_chain_time" binding:"required"`
}

// callback response code
const (
	SuccCode = "SUCCESS"
	FailCode = "FAIL"

	SuccMsg = "deal done"
)

// CallBackResponse ...
type CallBackResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}
