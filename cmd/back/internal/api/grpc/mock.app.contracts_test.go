// Code generated by MockGen. DO NOT EDIT.
// Source: grpc.go
//
// Generated by this command:
//
//	mockgen -source=grpc.go -destination mock.app.contracts_test.go -package grpc_test
//

// Package grpc_test is a generated GoMock package.
package grpc_test

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

// AddAvatar mocks base method.
func (m *Mockapplication) AddAvatar(ctx context.Context, session app.Session, fileID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAvatar", ctx, session, fileID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAvatar indicates an expected call of AddAvatar.
func (mr *MockapplicationMockRecorder) AddAvatar(ctx, session, fileID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAvatar", reflect.TypeOf((*Mockapplication)(nil).AddAvatar), ctx, session, fileID)
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

// CreateUser mocks base method.
func (m *Mockapplication) CreateUser(ctx context.Context, email, username, fullName, password string) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, email, username, fullName, password)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockapplicationMockRecorder) CreateUser(ctx, email, username, fullName, password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*Mockapplication)(nil).CreateUser), ctx, email, username, fullName, password)
}

// GetUsersByIDs mocks base method.
func (m *Mockapplication) GetUsersByIDs(ctx context.Context, session app.Session, ids []uuid.UUID) ([]app.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersByIDs", ctx, session, ids)
	ret0, _ := ret[0].([]app.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersByIDs indicates an expected call of GetUsersByIDs.
func (mr *MockapplicationMockRecorder) GetUsersByIDs(ctx, session, ids any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersByIDs", reflect.TypeOf((*Mockapplication)(nil).GetUsersByIDs), ctx, session, ids)
}

// ListUserAvatars mocks base method.
func (m *Mockapplication) ListUserAvatars(ctx context.Context, session app.Session) ([]app.AvatarInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUserAvatars", ctx, session)
	ret0, _ := ret[0].([]app.AvatarInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUserAvatars indicates an expected call of ListUserAvatars.
func (mr *MockapplicationMockRecorder) ListUserAvatars(ctx, session any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserAvatars", reflect.TypeOf((*Mockapplication)(nil).ListUserAvatars), ctx, session)
}

// ListUserByFilters mocks base method.
func (m *Mockapplication) ListUserByFilters(ctx context.Context, arg1 app.Session, filters app.SearchParams) ([]app.User, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUserByFilters", ctx, arg1, filters)
	ret0, _ := ret[0].([]app.User)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListUserByFilters indicates an expected call of ListUserByFilters.
func (mr *MockapplicationMockRecorder) ListUserByFilters(ctx, arg1, filters any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserByFilters", reflect.TypeOf((*Mockapplication)(nil).ListUserByFilters), ctx, arg1, filters)
}

// Login mocks base method.
func (m *Mockapplication) Login(ctx context.Context, email, password string, origin app.Origin) (uuid.UUID, *app.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, email, password, origin)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(*app.Token)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Login indicates an expected call of Login.
func (mr *MockapplicationMockRecorder) Login(ctx, email, password, origin any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*Mockapplication)(nil).Login), ctx, email, password, origin)
}

// Logout mocks base method.
func (m *Mockapplication) Logout(ctx context.Context, session app.Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout", ctx, session)
	ret0, _ := ret[0].(error)
	return ret0
}

// Logout indicates an expected call of Logout.
func (mr *MockapplicationMockRecorder) Logout(ctx, session any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*Mockapplication)(nil).Logout), ctx, session)
}

// RemoveAvatar mocks base method.
func (m *Mockapplication) RemoveAvatar(ctx context.Context, session app.Session, fileID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAvatar", ctx, session, fileID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAvatar indicates an expected call of RemoveAvatar.
func (mr *MockapplicationMockRecorder) RemoveAvatar(ctx, session, fileID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAvatar", reflect.TypeOf((*Mockapplication)(nil).RemoveAvatar), ctx, session, fileID)
}

// UpdatePassword mocks base method.
func (m *Mockapplication) UpdatePassword(ctx context.Context, session app.Session, oldPass, newPass string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", ctx, session, oldPass, newPass)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePassword indicates an expected call of UpdatePassword.
func (mr *MockapplicationMockRecorder) UpdatePassword(ctx, session, oldPass, newPass any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*Mockapplication)(nil).UpdatePassword), ctx, session, oldPass, newPass)
}

// UpdateUser mocks base method.
func (m *Mockapplication) UpdateUser(ctx context.Context, session app.Session, username string, avatarID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, session, username, avatarID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockapplicationMockRecorder) UpdateUser(ctx, session, username, avatarID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*Mockapplication)(nil).UpdateUser), ctx, session, username, avatarID)
}

// UserByID mocks base method.
func (m *Mockapplication) UserByID(ctx context.Context, session app.Session, userID uuid.UUID) (*app.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserByID", ctx, session, userID)
	ret0, _ := ret[0].(*app.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserByID indicates an expected call of UserByID.
func (mr *MockapplicationMockRecorder) UserByID(ctx, session, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserByID", reflect.TypeOf((*Mockapplication)(nil).UserByID), ctx, session, userID)
}

// VerificationEmail mocks base method.
func (m *Mockapplication) VerificationEmail(ctx context.Context, email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerificationEmail", ctx, email)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerificationEmail indicates an expected call of VerificationEmail.
func (mr *MockapplicationMockRecorder) VerificationEmail(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerificationEmail", reflect.TypeOf((*Mockapplication)(nil).VerificationEmail), ctx, email)
}

// VerificationUsername mocks base method.
func (m *Mockapplication) VerificationUsername(ctx context.Context, username string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerificationUsername", ctx, username)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerificationUsername indicates an expected call of VerificationUsername.
func (mr *MockapplicationMockRecorder) VerificationUsername(ctx, username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerificationUsername", reflect.TypeOf((*Mockapplication)(nil).VerificationUsername), ctx, username)
}
