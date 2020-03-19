/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package callback

//go:generate mockgen -destination=components/callback/mock/mock_callback.go -package=mock_callback github.com/csiabb/blockchain-adapter/components/callback ICallback

// ICallback callback interface
type ICallback interface {
	PushMessage(msg []byte, url string) ([]byte, error)
}
