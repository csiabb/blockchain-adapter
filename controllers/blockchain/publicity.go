/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package blockchain

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/csiabb/blockchain-adapter/common/rest"
	"github.com/csiabb/blockchain-adapter/structs/api"

	"github.com/gin-gonic/gin"
)

// PublicityData addition publicity data to blockchain
// @param 'uid' {string} account id on blockchain system
// @param 'publicity' {string} the json string of publicity data
// @result 'id' {string} the id of blockchain system
func (h *RestHandler) PublicityData(c *gin.Context) {
	logger.Info("got publicity data request")

	req := &api.PublicityRequest{}
	err := c.BindJSON(req)
	if nil != err {
		logger.Errorf("parse request params error: %v", err)
		c.JSON(http.StatusBadRequest, rest.ErrorResponse(rest.ParamsParseErr, err.Error()))
		return
	}
	logger.Infof("request params: %+v", req)

	var data interface{}
	err = json.Unmarshal([]byte(req.Publicity), &data)
	if nil != err {
		logger.Errorf("parse publicity by json error: %v", err)
		c.JSON(http.StatusBadRequest, rest.ErrorResponse(rest.ParamsInvalidErr,
			fmt.Sprintf("param publicity invalid, %s", err.Error())))
		return
	}

	c.JSON(http.StatusOK, rest.SuccessResponse(&api.BlockchainResponse{ID: ""}))
	logger.Infof("response publicity data success.")
	return
}
