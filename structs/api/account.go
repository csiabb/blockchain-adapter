/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package api

// CreateAccountRequest define create account request
type CreateAccountRequest struct {
	AccountID string `json:"account_id" binding:"required"`
}
