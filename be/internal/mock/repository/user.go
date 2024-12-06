// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/repository/user.go
//
// Generated by this command:
//
//	mockgen -destination internal/mock/repository/user.go -package=repository_mock -source=internal/app/repository/user.go UserRepository
//

// Package repository_mock is a generated GoMock package.
package repository_mock

import (
	context "context"
	reflect "reflect"

	persistence "github.com/nnrmps/blue-vending-machine/be/internal/app/persistence"
	gomock "go.uber.org/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
	isgomock struct{}
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockUserRepository) Login(ctx context.Context, tx *gorm.DB, userName, password string) (persistence.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, tx, userName, password)
	ret0, _ := ret[0].(persistence.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockUserRepositoryMockRecorder) Login(ctx, tx, userName, password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserRepository)(nil).Login), ctx, tx, userName, password)
}