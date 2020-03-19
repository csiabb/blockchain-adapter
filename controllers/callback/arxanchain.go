/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package callback

import (
	"encoding/json"
	"net/http"

	arxan "github.com/csiabb/blockchain-adapter/adapter/arxanchain/structs"
	"github.com/csiabb/blockchain-adapter/structs/api"

	"github.com/gin-gonic/gin"
)

// ArxanchainCallback arxanchain callback handler
func (h *RestHandler) ArxanchainCallback(c *gin.Context) {
	logger.Infof("Got arxanchain callback notification")

	req := &arxan.CallBackRequest{}
	err := c.BindJSON(req)
	if nil != err {
		logger.Errorf("parse request params error: %v", err)
		c.JSON(http.StatusBadRequest, &arxan.CallBackResponse{
			Code: arxan.FailCode,
			Msg:  err.Error(),
		})
		return
	}
	logger.Infof("request params: %+v", req)

	c.JSON(http.StatusOK, &arxan.CallBackResponse{
		Code: arxan.SuccCode,
		Msg:  arxan.SuccMsg,
	})

	logger.Infof("Response arxanchain blockchain callback notification success.")

	notification := &api.CallbackNotification{
		Blockchain: arxan.BlockchainName,
		ID:         req.ID,
		BlockNum:   int64(req.BlockNumber),
		TxID:       req.TxID,
		Time:       req.OnChainTime,
	}
	pushBytes, err := json.Marshal(notification)
	if nil != err {
		logger.Errorf("Marshal notification content error, %v", err)
		return
	}

	pushResBytes, err := h.srvcContext.CallbackClient.PushMessage(pushBytes, h.srvcContext.Config.Callback.URL)
	if nil != err {
		logger.Errorf("push arxanchain notification %s error, %v", req.ID, err)
		return
	}

	var respBody api.CallbackResponse
	err = json.Unmarshal(pushResBytes, &respBody)
	if err != nil {
		logger.Errorf("Parse push arxanchain notification %s result error: %v", req.ID, err)
		return
	}
	logger.Infof("push arxanchain notification %s result %v", req.ID, respBody)

	return
}
