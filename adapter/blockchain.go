/*
Copyright Arxan Chain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package adapter

import (
	"blockchain-adapter/structs"
)

// BlockchainAdapter adapter of blockchain
type BlockchainAdapter interface {
	CommitTransaction(*structs.BCTransactionData) interface{}
}
