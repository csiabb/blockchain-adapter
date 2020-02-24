/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package structs

// CommonRequest common rest request
type CommonRequest struct {
	Action  string
	Request interface{}
}

// CommonResponse common rest response
type CommonResponse struct {
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

// QueryAccountRequest ...
type QueryAccountRequest struct {
	Phone string
	Email string
	UID   string
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

// QueryAssetDetailRequest ...
type QueryAssetDetailRequest struct {
	AssetDID string
}

// QueryAssetDetailResponse ...
type QueryAssetDetailResponse struct {
	AssetDID   string   `json:"asset_did"`  // 数据资产DID
	AssetName  string   `json:"asset_name"` // 数据资产名称
	OwnerDID   string   `json:"owner_did"`  // 数据资产所属人
	Intro      string   `json:"intro"`      // 资产简介
	Category   string   `json:"category"`   // 资产类型
	Visibility string   `json:"visibility"` // 可见性
	Tags       []string `json:"tags"`       // 数据资产标签列表
	Metadata   string   `json:"metadata"`   // 数据资产数据
	TxID       string   `json:"tx_id"`      // 数据资产上链交易ID
	CreatedAt  int64    `json:"created_at"` // 数据资产创建登记时间
}
