/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/csiabb/blockchain-adapter/adapter/arxanchain/config"
)

// ArxanchainClient aranchain block client
type ArxanchainClient struct {
	config     *config.Config
	httpClient *http.Client
}

// NewArxanchainClient new arxanchain client
func NewArxanchainClient(cfg *config.Config) (*ArxanchainClient, error) {
	if nil == cfg {
		return nil, fmt.Errorf("config is nil")
	}

	if "" == cfg.Host {
		return nil, fmt.Errorf("Missing host config")
	}

	if "" == cfg.APIKey {
		return nil, fmt.Errorf("Missing api key config")
	}

	if "" == cfg.APISecret {
		return nil, fmt.Errorf("Missing api secret config")
	}

	parts := strings.SplitN(cfg.Host, "://", 2)
	if len(parts) == 2 {
		switch parts[0] {
		case "http":
			cfg.Scheme = "http"
		case "https":
			cfg.Scheme = "https"
		default:
			return nil, fmt.Errorf("Unknown protocol scheme: %s", parts[0])
		}
		cfg.Host = parts[1]
	}

	c := &ArxanchainClient{
		config:     cfg,
		httpClient: &http.Client{},
	}

	return c, nil
}

// NewRequest is used to create a new request
func (c *ArxanchainClient) NewRequest(method, path string) *Request {
	r := &Request{
		method: method,
		url: &url.URL{
			Scheme: c.config.Scheme,
			Host:   c.config.Host,
			Path:   path,
		},
		params: make(map[string][]string),
		header: make(http.Header),
	}
	// Prevent data being compressed by gzip
	r.header.Set("Accept-Encoding", "*")
	return r
}

// DoRequest does an HTTP request.
//
// Once the crypto mode enabled, the response result will
// be decrypted and verified signature using crypto libary,
// then the final result will be return to end client.
//
func (c *ArxanchainClient) DoRequest(r *Request) (time.Duration, *http.Response, error) {
	req, err := r.ToHTTP()
	if err != nil {
		return 0, nil, err
	}
	start := time.Now()

	resp, err := c.HTTPClient.Do(req)
	diff := time.Since(start)

	// decrypt and verify resp.Body
	if err != nil {
		return diff, resp, err
	}
	if resp == nil {
		return diff, nil, fmt.Errorf("DoRequest http.Response is nil")
	}

	return diff, resp, err
}

// DecodeBody is used to JSON decode a body
func DecodeBody(resp *http.Response, out interface{}) error {
	dec := json.NewDecoder(resp.Body)
	return dec.Decode(out)
}

// EncodeBody is used to encode a request body
func EncodeBody(obj interface{}) (io.Reader, error) {
	buf := bytes.NewBuffer(nil)
	enc := json.NewEncoder(buf)
	if err := enc.Encode(obj); err != nil {
		return nil, err
	}
	return buf, nil
}

// RequireOK is used to wrap DoRequest
func RequireOK(d time.Duration, resp *http.Response, e error) (time.Duration, *http.Response, error) {
	if e != nil {
		if resp != nil {
			resp.Body.Close()
		}
		return d, nil, e
	}
	return d, resp, nil
}
