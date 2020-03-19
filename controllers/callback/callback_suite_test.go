/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package callback_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCallback(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Callback Suite")
}
