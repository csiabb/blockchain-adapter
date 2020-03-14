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

var _ = Describe("publicity data api test", func() {
	var (
		host        = "http://localhost:8989"
		path        = "/api/v1/blockchain/publicities"
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

	Describe("PublicityData api test", func() {
		Context("request params error", func() {
			It("missing required param uid and publicity", func() {
				paramStr := `{}`
				bodyBuf := bytes.NewBufferString(paramStr)

				target := host + path

				resp := httptest.NewRecorder()
				c, router := gin.CreateTestContext(resp)
				router.POST(path, handler.PublicityData)
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
			It("param publicity not json string", func() {
				paramStr := `{"uid":"uniteid","publicity":"not json string"}`
				bodyBuf := bytes.NewBufferString(paramStr)

				target := host + path

				resp := httptest.NewRecorder()
				c, router := gin.CreateTestContext(resp)
				router.POST(path, handler.PublicityData)
				c.Request = httptest.NewRequest(http.MethodPost, target, bodyBuf)

				router.ServeHTTP(resp, c.Request)

				result := resp.Result()
				defer result.Body.Close()
				body, _ := ioutil.ReadAll(result.Body)
				response := rest.CommonResponse{}
				_ = json.Unmarshal(body, &response)

				Expect(resp.Code).Should(Equal(http.StatusBadRequest))
				Expect(int(response.Code)).Should(Equal(rest.ParamsInvalidErr))
			})
		})
		Context("request handler error", func() {
			It("arxanchain publicity data error", func() {
				paramStr := `{"uid":"uniteid","publicity":"{}"}`
				bodyBuf := bytes.NewBufferString(paramStr)

				target := host + path

				resp := httptest.NewRecorder()
				c, router := gin.CreateTestContext(resp)
				router.POST(path, handler.PublicityData)
				c.Request = httptest.NewRequest(http.MethodPost, target, bodyBuf)

				mockAdapter.EXPECT().PublicityData(gomock.Any()).Return(nil, fmt.Errorf("arxanchain publicity data error"))

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
			It("publicity data success", func() {
				paramStr := `{"uid":"uniteid","publicity":"{}"}`
				bodyBuf := bytes.NewBufferString(paramStr)

				target := host + path

				resp := httptest.NewRecorder()
				c, router := gin.CreateTestContext(resp)
				router.POST(path, handler.PublicityData)
				c.Request = httptest.NewRequest(http.MethodPost, target, bodyBuf)

				mockAdapter.EXPECT().PublicityData(gomock.Any()).Return(&adapter.BlockchainResponse{ID: "id"}, nil)

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
