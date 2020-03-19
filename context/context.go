/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package context

import (
	"fmt"

	"github.com/csiabb/blockchain-adapter/adapter"
	arxan "github.com/csiabb/blockchain-adapter/adapter/arxanchain/api"
	"github.com/csiabb/blockchain-adapter/common/log"
	"github.com/csiabb/blockchain-adapter/components/callback"
	"github.com/csiabb/blockchain-adapter/config"
)

var (
	serverContext *Context
	logger        = log.MustGetLogger("context")
)

// Context the context of service
type Context struct {
	Config           *config.SrvcCfg
	CallbackClient   callback.ICallback
	ArxanchainClient adapter.BlockchainAdapter
}

// GetServerContext ...
func GetServerContext() *Context {
	if serverContext == nil {
		serverContext = &Context{}
	}
	return serverContext
}

// Init init service context
func (c *Context) Init() error {
	if nil == c.Config {
		logger.Errorf("Initalize faild, configure is nil")
		return fmt.Errorf("configure is nil")
	}
	fmt.Println("init config:", c.Config)
	logger.Debugf("Initalization configure: %v", c.Config)

	var err error

	c.CallbackClient, err = callback.NewCallbackClient(&c.Config.Callback)
	if nil != err {
		logger.Errorf("new callback client error, %v", err)
		return fmt.Errorf("new callback client error: %s", err.Error())
	}

	if c.Config.Arxanchain.Enabled {
		c.ArxanchainClient, err = arxan.NewArxanchainClient(&c.Config.Arxanchain)
		if nil != err {
			logger.Errorf("new arxanchain client error, %v", err)
			return fmt.Errorf("new arxanchain client error: %s", err.Error())
		}
		logger.Infof("new arxanchain client success")
	}

	logger.Infof("initalize context success.")

	return nil
}
