/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package api

// PublicityRequest define publicity request
type PublicityRequest struct {
	UID       string `json:"uid" binding:"required"`
	Publicity string `json:"publicity" binding:"required"`
}
