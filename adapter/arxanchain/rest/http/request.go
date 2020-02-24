/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// Request is used to help build up a request
type Request struct {
	method string
	url    *url.URL
	params url.Values
	body   io.Reader
	header http.Header
	obj    interface{}
}

// SetBody is used to set Request body.
func (r *Request) SetBody(obj interface{}) error {
	if r.body != nil {
		return nil
	}
	if obj == nil {
		log.Println("Body object is nil")
		return fmt.Errorf("body object is nil")
	}

	// Check if we should encode the body
	var b io.Reader
	var objData []byte
	var err error
	var ok bool
	objData, ok = obj.([]byte)
	if ok && objData != nil {
		b = bytes.NewBuffer(objData)
	} else {
		b, err = EncodeBody(obj)
		if err != nil {
			log.Printf("EncodeBody fail: %v", err)
			return err
		}
	}

	r.body = b

	return nil
}

// SetHeaders is used to add multiple header KV pairs
func (r *Request) SetHeaders(headers http.Header) {
	for k, list := range headers {
		for _, v := range list {
			r.SetHeader(k, v)
		}
	}
}

// SetHeader is used to set one header KV pair
func (r *Request) SetHeader(k, v string) {
	if strings.ToLower(k) == "accept-encoding" {
		return
	}
	r.header.Set(k, v)
}

// GetHeader is used to get one header value
func (r *Request) GetHeader(k string) string {
	return r.header.Get(k)
}

// SetParamsWithMap is used to set multiple query params
func (r *Request) SetParamsWithMap(params map[string]string) {
	for k, v := range params {
		r.params.Set(k, v)
	}
}

// SetParams is used to set multiple query params
func (r *Request) SetParams(params url.Values) {
	for k, list := range params {
		for _, v := range list {
			r.params.Set(k, v)
		}
	}
}

// SetParam is used to set one query param
func (r *Request) SetParam(k, v string) {
	r.params.Set(k, v)
}

// GetParam is used to get one query param
func (r *Request) GetParam(k string) string {
	return r.params.Get(k)
}

// ToHTTP is used to convert the Request object to an standard HTTP request object.
func (r *Request) ToHTTP() (*http.Request, error) {
	// Encode the query parameters
	r.url.RawQuery = r.params.Encode()

	// Create the HTTP request
	req, err := http.NewRequest(r.method, r.url.RequestURI(), r.body)
	if err != nil {
		log.Printf("New http request fail: %v", err)
		return nil, err
	}

	req.URL.Host = r.url.Host
	req.URL.Scheme = r.url.Scheme
	req.Host = r.url.Host
	req.Header = r.header

	return req, nil
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
