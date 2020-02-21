/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package adapter

// BlockchainAdapter adapter of blockchain
type BlockchainAdapter interface {
	CommitTransaction(interface{}) (interface{}, error)
	QueryTransaction(interface{}) (interface{}, error)
}
