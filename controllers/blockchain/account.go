/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package blockchain

import (
	"net/http"

	"github.com/csiabb/blockchain-adapter/common/rest"
	"github.com/csiabb/blockchain-adapter/structs/api"

	"github.com/gin-gonic/gin"
)

// CreateAccount addition account information to blockchain
// @param 'account_id' {string} the account id in application system
// @result 'id' {string} the id of blockchain system
func (h *RestHandler) CreateAccount(c *gin.Context) {
	logger.Info("got create account request")

	req := &api.CreateAccountRequest{}
	err := c.BindJSON(req)
	if nil != err {
		logger.Errorf("parse request params error: %v", err)
		c.JSON(http.StatusBadRequest, rest.ErrorResponse(rest.ParamsParseErr, err.Error()))
		return
	}
	logger.Infof("request params: %+v", req)

	c.JSON(http.StatusOK, rest.SuccessResponse(&api.BlockchainResponse{ID: ""}))
	logger.Infof("response create account success.")
	return
}
