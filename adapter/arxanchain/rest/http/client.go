/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package http

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Client provides a client to the Rest API
type Client struct {
	HTTPClient *http.Client
	Cfg        *Config
}

// NewClient returns a new client
func NewClient(cfg *Config) (*Client, error) {
	if nil == cfg {
		return nil, fmt.Errorf("config is nil")
	}

	parts := strings.SplitN(cfg.Endpoint, "://", 2)
	if len(parts) == 2 {
		switch parts[0] {
		case "http":
			cfg.Scheme = "http"
		case "https":
			cfg.Scheme = "https"
		default:
			return nil, fmt.Errorf("Unknown protocol scheme: %s", parts[0])
		}
		cfg.Endpoint = parts[1]
	}

	return &Client{
		HTTPClient: &http.Client{},
		Cfg:        cfg,
	}, nil
}

// NewRequest is used to create a new request
func (c *Client) NewRequest(method, path string) *Request {
	r := &Request{
		method: method,
		url: &url.URL{
			Scheme: c.Cfg.Scheme,
			Host:   c.Cfg.Endpoint,
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
func (c *Client) DoRequest(r *Request) (time.Duration, *http.Response, error) {
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
