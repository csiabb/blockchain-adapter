/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package callback_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	arxan "github.com/csiabb/blockchain-adapter/adapter/arxanchain/structs"
	push "github.com/csiabb/blockchain-adapter/components/callback"
	mockpush "github.com/csiabb/blockchain-adapter/components/callback/mock"
	"github.com/csiabb/blockchain-adapter/config"
	"github.com/csiabb/blockchain-adapter/context"
	"github.com/csiabb/blockchain-adapter/controllers/callback"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Version api test", func() {
	var (
		host     = "http://localhost:8989"
		path     = "/api/v1/blockchain/callback/arxan"
		handler  *callback.RestHandler
		ctrl     *gomock.Controller
		mockPush *mockpush.MockICallback
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		ctrl = gomock.NewController(GinkgoT())

		mockPush = mockpush.NewMockICallback(ctrl)
		h, err := callback.NewRestHandler(&context.Context{
			Config: &config.SrvcCfg{
				Callback: push.Config{
					TimeOut: 5,
					URL:     "https://csiabb.com/callback",
				},
			},
			CallbackClient: mockPush,
		})

		Expect(err).Should(BeNil())
		handler = h
	})

	AfterEach(func() {
		handler = nil
		ctrl.Finish()
	})

	Describe("Arxanchain callback api test", func() {
		Context("arxanchain callback error", func() {
			It("arxanchain callback notification missing required param", func() {
				paramStr := `{"id":"id","type":"DigitalAsset","status":"SUCCESS","tx_id":"txid","on_chain_time":1583826856}`
				bodyBuf := bytes.NewBufferString(paramStr)

				target := host + path

				resp := httptest.NewRecorder()
				c, router := gin.CreateTestContext(resp)
				router.POST(path, handler.ArxanchainCallback)
				c.Request = httptest.NewRequest(http.MethodPost, target, bodyBuf)

				router.ServeHTTP(resp, c.Request)

				result := resp.Result()
				defer result.Body.Close()
				body, _ := ioutil.ReadAll(result.Body)
				response := arxan.CallBackResponse{}
				_ = json.Unmarshal(body, &response)

				Expect(resp.Code).Should(Equal(http.StatusBadRequest))
				Expect(response.Code).Should(Equal(arxan.FailCode))
			})
		})
		Context("request success", func() {
			It("Arxanchain callback notification handle success", func() {
				paramStr := `{"id":"id","type":"DigitalAsset","status":"SUCCESS","block_number":123,"tx_id":"txid","on_chain_time":1583826856}`
				bodyBuf := bytes.NewBufferString(paramStr)

				target := host + path

				resp := httptest.NewRecorder()
				c, router := gin.CreateTestContext(resp)
				router.POST(path, handler.ArxanchainCallback)
				c.Request = httptest.NewRequest(http.MethodPost, target, bodyBuf)

				pushResult := []byte(`{"code":"success","msg":""}`)

				mockPush.EXPECT().PushMessage(gomock.Any(), gomock.Any()).Return(pushResult, nil)

				router.ServeHTTP(resp, c.Request)

				result := resp.Result()
				defer result.Body.Close()
				body, _ := ioutil.ReadAll(result.Body)
				response := arxan.CallBackResponse{}
				_ = json.Unmarshal(body, &response)

				Expect(resp.Code).Should(Equal(http.StatusOK))
				Expect(response.Code).Should(Equal(arxan.SuccCode))
				Expect(response.Msg).Should(Equal(arxan.SuccMsg))
			})
		})
	})
})
