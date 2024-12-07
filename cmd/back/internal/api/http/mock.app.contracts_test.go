// Code generated by MockGen. DO NOT EDIT.
// Source: http.go
//
// Generated by this command:
//
//	mockgen -source=http.go -destination mock.app.contracts_test.go -package http_test
//

// Package http_test is a generated GoMock package.
package http_test

import (
	context "context"
	reflect "reflect"

	app "github.com/ZergsLaw/back-template/cmd/back/internal/app"
	uuid "github.com/gofrs/uuid"
	gomock "go.uber.org/mock/gomock"
)

// Mockapplication is a mock of application interface.
type Mockapplication struct {
	ctrl     *gomock.Controller
	recorder *MockapplicationMockRecorder
	isgomock struct{}
}

// MockapplicationMockRecorder is the mock recorder for Mockapplication.
type MockapplicationMockRecorder struct {
	mock *Mockapplication
}

// NewMockapplication creates a new mock instance.
func NewMockapplication(ctrl *gomock.Controller) *Mockapplication {
	mock := &Mockapplication{ctrl: ctrl}
	mock.recorder = &MockapplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockapplication) EXPECT() *MockapplicationMockRecorder {
	return m.recorder
}

// Auth mocks base method.
func (m *Mockapplication) Auth(ctx context.Context, token string) (*app.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Auth", ctx, token)
	ret0, _ := ret[0].(*app.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Auth indicates an expected call of Auth.
func (mr *MockapplicationMockRecorder) Auth(ctx, token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Auth", reflect.TypeOf((*Mockapplication)(nil).Auth), ctx, token)
}

// GetFile mocks base method.
func (m *Mockapplication) GetFile(ctx context.Context, session app.Session, fileID uuid.UUID) (*app.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFile", ctx, session, fileID)
	ret0, _ := ret[0].(*app.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFile indicates an expected call of GetFile.
func (mr *MockapplicationMockRecorder) GetFile(ctx, session, fileID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFile", reflect.TypeOf((*Mockapplication)(nil).GetFile), ctx, session, fileID)
}

// SaveFile mocks base method.
func (m *Mockapplication) SaveFile(ctx context.Context, session app.Session, file app.File) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveFile", ctx, session, file)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveFile indicates an expected call of SaveFile.
func (mr *MockapplicationMockRecorder) SaveFile(ctx, session, file any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveFile", reflect.TypeOf((*Mockapplication)(nil).SaveFile), ctx, session, file)
}
