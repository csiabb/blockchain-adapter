/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package blockchain_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/csiabb/blockchain-adapter/adapter"
	arxan "github.com/csiabb/blockchain-adapter/adapter/arxanchain/rest/http"
	mockadapter "github.com/csiabb/blockchain-adapter/adapter/mock"
	"github.com/csiabb/blockchain-adapter/common/rest"
	"github.com/csiabb/blockchain-adapter/config"
	"github.com/csiabb/blockchain-adapter/context"
	"github.com/csiabb/blockchain-adapter/controllers/blockchain"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Account api test", func() {
	var (
		host        = "http://localhost:8989"
		path        = "/api/v1/blockchain/accounts"
		handler     *blockchain.RestHandler
		ctrl        *gomock.Controller
		mockAdapter *mockadapter.MockBlockchainAdapter
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		ctrl = gomock.NewController(GinkgoT())

		mockAdapter = mockadapter.NewMockBlockchainAdapter(ctrl)
		ctx := &context.Context{
			Config: &config.SrvcCfg{
				Arxanchain: arxan.Config{
					Enabled: true,
				},
			},
			ArxanchainClient: mockAdapter,
		}
		h, err := blockchain.NewRestHandler(ctx)
		Expect(err).Should(BeNil())
		handler = h
	})

	AfterEach(func() {
		handler = nil
		ctrl.Finish()
	})

	Describe("CreateAccount api test", func() {
		Context("request params error", func() {
			It("missing required param account_id", func() {
				paramStr := `{}`
				bodyBuf := bytes.NewBufferString(paramStr)

				target := host + path

				resp := httptest.NewRecorder()
				c, router := gin.CreateTestContext(resp)
				router.POST(path, handler.CreateAccount)
				c.Request = httptest.NewRequest(http.MethodPost, target, bodyBuf)

				router.ServeHTTP(resp, c.Request)

				result := resp.Result()
				defer result.Body.Close()
				body, _ := ioutil.ReadAll(result.Body)
				response := rest.CommonResponse{}
				_ = json.Unmarshal(body, &response)

				Expect(resp.Code).Should(Equal(http.StatusBadRequest))
				Expect(int(response.Code)).Should(Equal(rest.ParamsParseErr))
			})
		})
		Context("request handler error", func() {
			It("arxanchain create error", func() {
				paramStr := `{"account_id":"uniteid"}`
				bodyBuf := bytes.NewBufferString(paramStr)

				target := host + path

				resp := httptest.NewRecorder()
				c, router := gin.CreateTestContext(resp)
				router.POST(path, handler.CreateAccount)
				c.Request = httptest.NewRequest(http.MethodPost, target, bodyBuf)

				mockAdapter.EXPECT().CreateAccount(gomock.Any()).Return(nil, fmt.Errorf("arxanchain create error"))

				router.ServeHTTP(resp, c.Request)

				result := resp.Result()
				defer result.Body.Close()
				body, _ := ioutil.ReadAll(result.Body)
				response := rest.CommonResponse{}
				_ = json.Unmarshal(body, &response)

				Expect(resp.Code).Should(Equal(http.StatusInternalServerError))
				Expect(int(response.Code)).Should(Equal(rest.InternalServerErr))
			})
		})
		Context("request success", func() {
			It("create account success", func() {
				paramStr := `{"account_id":"uniteid"}`
				bodyBuf := bytes.NewBufferString(paramStr)

				target := host + path

				resp := httptest.NewRecorder()
				c, router := gin.CreateTestContext(resp)
				router.POST(path, handler.CreateAccount)
				c.Request = httptest.NewRequest(http.MethodPost, target, bodyBuf)

				mockAdapter.EXPECT().CreateAccount(gomock.Any()).Return(&adapter.BlockchainResponse{ID: "id"}, nil)

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
