/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package callback

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/csiabb/blockchain-adapter/common/log"
)

var (
	logger = log.MustGetLogger("component-callback")
)

// callback client const var
const (
	defaultPushTimeOut = 5 // 5s
)

// Config callback configure
type Config struct {
	TimeOut int
	URL     string
}

// Client ...
type Client struct {
	cfg        *Config
	pushClient *http.Client
}

// NewCallbackClient new callback client
func NewCallbackClient(c *Config) (*Client, error) {
	if nil == c {
		return nil, fmt.Errorf("configure is nil")
	}
	if 0 >= c.TimeOut {
		c.TimeOut = defaultPushTimeOut
	}
	timeOut := time.Duration(c.TimeOut) * time.Second
	client := &Client{
		cfg: c,
		pushClient: &http.Client{
			Transport: &http.Transport{
				Dial: func(netw, addr string) (net.Conn, error) {
					conn, err := net.DialTimeout(netw, addr, timeOut)
					if err != nil {
						return nil, err
					}
					conn.SetDeadline(time.Now().Add(timeOut))
					return conn, nil
				},
				ResponseHeaderTimeout: timeOut,
			},
		},
	}
	return client, nil
}

// PushMessage ...
// TODO: use message queue or channel
func (c *Client) PushMessage(msg []byte, url string) ([]byte, error) {
	body := bytes.NewBuffer(msg)
	resp, err := c.pushClient.Post(url, "application/json;charset=utf-8", body)
	if err != nil {
		logger.Errorf("push message to %s error, %v", url, err)
		return nil, err
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Errorf("read push message result error, %v", err)
		return nil, err
	}
	return result, nil
}
