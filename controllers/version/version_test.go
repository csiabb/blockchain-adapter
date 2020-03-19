/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package version_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/csiabb/blockchain-adapter/common/rest"
	"github.com/csiabb/blockchain-adapter/context"
	"github.com/csiabb/blockchain-adapter/controllers/version"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Version api test", func() {
	var (
		host    = "http://localhost:8989"
		path    = "/verison"
		handler *version.RestHandler
		ctrl    *gomock.Controller
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		ctrl = gomock.NewController(GinkgoT())

		h, err := version.NewRestHandler(&context.Context{})
		Expect(err).Should(BeNil())
		handler = h
	})

	AfterEach(func() {
		handler = nil
		ctrl.Finish()
	})

	Describe("Version api test", func() {
		Context("request success", func() {
			It("Query service version success", func() {
				target := host + path

				resp := httptest.NewRecorder()
				c, router := gin.CreateTestContext(resp)
				router.GET(path, handler.Version)
				c.Request = httptest.NewRequest(http.MethodGet, target, nil)

				router.ServeHTTP(resp, c.Request)

				result := resp.Result()
				defer result.Body.Close()
				body, _ := ioutil.ReadAll(result.Body)
				response := rest.CommonResponse{}
				_ = json.Unmarshal(body, &response)

				Expect(resp.Code).Should(Equal(http.StatusOK))
				Expect(int(response.Code)).Should(Equal(rest.SuccCode))
			})
		})
	})
})
