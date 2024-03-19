// Code generated by MockGen. DO NOT EDIT.
// Source: internal/server/repository/keeper.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	repository "github.com/Galish/goph-keeper/internal/server/repository"
	gomock "github.com/golang/mock/gomock"
)

// MockKeeperRepository is a mock of KeeperRepository interface.
type MockKeeperRepository struct {
	ctrl     *gomock.Controller
	recorder *MockKeeperRepositoryMockRecorder
}

// MockKeeperRepositoryMockRecorder is the mock recorder for MockKeeperRepository.
type MockKeeperRepositoryMockRecorder struct {
	mock *MockKeeperRepository
}

// NewMockKeeperRepository creates a new mock instance.
func NewMockKeeperRepository(ctrl *gomock.Controller) *MockKeeperRepository {
	mock := &MockKeeperRepository{ctrl: ctrl}
	mock.recorder = &MockKeeperRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKeeperRepository) EXPECT() *MockKeeperRepositoryMockRecorder {
	return m.recorder
}

// CreateSecureRecord mocks base method.
func (m *MockKeeperRepository) CreateSecureRecord(arg0 context.Context, arg1 *repository.SecureRecord) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSecureRecord", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSecureRecord indicates an expected call of CreateSecureRecord.
func (mr *MockKeeperRepositoryMockRecorder) CreateSecureRecord(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecureRecord", reflect.TypeOf((*MockKeeperRepository)(nil).CreateSecureRecord), arg0, arg1)
}

// GetSecureRecord mocks base method.
func (m *MockKeeperRepository) GetSecureRecord(arg0 context.Context, arg1, arg2 string) (*repository.SecureRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecureRecord", arg0, arg1, arg2)
	ret0, _ := ret[0].(*repository.SecureRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecureRecord indicates an expected call of GetSecureRecord.
func (mr *MockKeeperRepositoryMockRecorder) GetSecureRecord(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecureRecord", reflect.TypeOf((*MockKeeperRepository)(nil).GetSecureRecord), arg0, arg1, arg2)
}

// GetSecureRecords mocks base method.
func (m *MockKeeperRepository) GetSecureRecords(arg0 context.Context, arg1 string, arg2 repository.SecureRecordType) ([]*repository.SecureRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecureRecords", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*repository.SecureRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecureRecords indicates an expected call of GetSecureRecords.
func (mr *MockKeeperRepositoryMockRecorder) GetSecureRecords(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecureRecords", reflect.TypeOf((*MockKeeperRepository)(nil).GetSecureRecords), arg0, arg1, arg2)
}
