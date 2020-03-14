/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package http

// Config arxanchain client config
type Config struct {
	Enabled   bool
	Name      string
	Scheme    string
	Endpoint  string // arxanchain service endpoint
	APIKey    string // api key
	APISecret string // api secret
	Callback  string // notify blockchain result callback url
}
