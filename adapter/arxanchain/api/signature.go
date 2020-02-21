/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/csiabb/blockchain-adapter/adapter/arxanchain/structs"
)

// SignatureData ...
type SignatureData struct {
	Secret        string // 签名秘钥
	Signature     string // 签名字符串
	Message       string // 字符串
	SignAlgo      string // HMAC-MD5, HMAC-SHA256
	RequestPath   string // 请求路径
	RequestMethod string // 请求方法
	Timestamp     int64  // 签名时间戳
}

// Signature ...
// sig data: [Parameter Message] + [Request Path] + [Request Method] + [timestamp]
// where:
// Parameter Message is the data being sent to your API.
// Request Path is an identifier that your API will use to recognise the service (remember that this won’t say who the service is acting on behalf of).
// Request Method is GET/POST etc.
// Timestamp corresponds to when the request was made.
func Signature(sig *SignatureData) (string, error) {
	// hash算法
	signAlgo := sig.SignAlgo
	if signAlgo == "" {
		signAlgo = structs.HMACSHA256
	}

	// 签名字符串
	data := fmt.Sprintf("%s%s%s%d", sig.Message, sig.RequestPath, sig.RequestMethod, sig.Timestamp)

	switch signAlgo {
	case structs.HMACSHA256:
		return getHmacSHA256Code(sig.Secret, data), nil
	default:
		return "", fmt.Errorf("%s", "Invalid sign algo")
	}
}

// getHmacSHA256Code ...
func getHmacSHA256Code(secret, message string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	// Get result and encode as hexadecimal string
	return hex.EncodeToString(h.Sum(nil))
}
