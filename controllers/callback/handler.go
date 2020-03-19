/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package callback

import (
	"github.com/csiabb/blockchain-adapter/common/log"
	"github.com/csiabb/blockchain-adapter/context"
)

var (
	logger = log.MustGetLogger("callback-handler")
)

// RestHandler callback handler
type RestHandler struct {
	srvcContext *context.Context
}

// NewRestHandler ...
func NewRestHandler(c *context.Context) (*RestHandler, error) {
	return &RestHandler{srvcContext: c}, nil
}
