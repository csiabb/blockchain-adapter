// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/csiabb/blockchain-adapter/adapter (interfaces: BlockchainAdapter)

// Package mock_adapter is a generated GoMock package.
package mock_adapter

import (
	adapter "github.com/csiabb/blockchain-adapter/adapter"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockBlockchainAdapter is a mock of BlockchainAdapter interface
type MockBlockchainAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockBlockchainAdapterMockRecorder
}

// MockBlockchainAdapterMockRecorder is the mock recorder for MockBlockchainAdapter
type MockBlockchainAdapterMockRecorder struct {
	mock *MockBlockchainAdapter
}

// NewMockBlockchainAdapter creates a new mock instance
func NewMockBlockchainAdapter(ctrl *gomock.Controller) *MockBlockchainAdapter {
	mock := &MockBlockchainAdapter{ctrl: ctrl}
	mock.recorder = &MockBlockchainAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBlockchainAdapter) EXPECT() *MockBlockchainAdapterMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method
func (m *MockBlockchainAdapter) CreateAccount(arg0 *adapter.CreateAccountReq) (*adapter.BlockchainResponse, error) {
	ret := m.ctrl.Call(m, "CreateAccount", arg0)
	ret0, _ := ret[0].(*adapter.BlockchainResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount
func (mr *MockBlockchainAdapterMockRecorder) CreateAccount(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockBlockchainAdapter)(nil).CreateAccount), arg0)
}

// PublicityData mocks base method
func (m *MockBlockchainAdapter) PublicityData(arg0 *adapter.PublicityDataReq) (*adapter.BlockchainResponse, error) {
	ret := m.ctrl.Call(m, "PublicityData", arg0)
	ret0, _ := ret[0].(*adapter.BlockchainResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublicityData indicates an expected call of PublicityData
func (mr *MockBlockchainAdapterMockRecorder) PublicityData(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicityData", reflect.TypeOf((*MockBlockchainAdapter)(nil).PublicityData), arg0)
}
