// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/TA-MoQ/quic-go (interfaces: ConnRunner)
//
// Generated by this command:
//
//	mockgen.exe -typed -build_flags=-tags=gomock -package quic -self_package github.com/TA-MoQ/quic-go -destination mock_conn_runner_test.go github.com/TA-MoQ/quic-go ConnRunner
//
// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	protocol "github.com/TA-MoQ/quic-go/internal/protocol"
	gomock "go.uber.org/mock/gomock"
)

// MockConnRunner is a mock of ConnRunner interface.
type MockConnRunner struct {
	ctrl     *gomock.Controller
	recorder *MockConnRunnerMockRecorder
}

// MockConnRunnerMockRecorder is the mock recorder for MockConnRunner.
type MockConnRunnerMockRecorder struct {
	mock *MockConnRunner
}

// NewMockConnRunner creates a new mock instance.
func NewMockConnRunner(ctrl *gomock.Controller) *MockConnRunner {
	mock := &MockConnRunner{ctrl: ctrl}
	mock.recorder = &MockConnRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnRunner) EXPECT() *MockConnRunnerMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockConnRunner) Add(arg0 protocol.ConnectionID, arg1 packetHandler) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockConnRunnerMockRecorder) Add(arg0, arg1 any) *ConnRunnerAddCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockConnRunner)(nil).Add), arg0, arg1)
	return &ConnRunnerAddCall{Call: call}
}

// ConnRunnerAddCall wrap *gomock.Call
type ConnRunnerAddCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ConnRunnerAddCall) Return(arg0 bool) *ConnRunnerAddCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ConnRunnerAddCall) Do(f func(protocol.ConnectionID, packetHandler) bool) *ConnRunnerAddCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ConnRunnerAddCall) DoAndReturn(f func(protocol.ConnectionID, packetHandler) bool) *ConnRunnerAddCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AddResetToken mocks base method.
func (m *MockConnRunner) AddResetToken(arg0 protocol.StatelessResetToken, arg1 packetHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddResetToken", arg0, arg1)
}

// AddResetToken indicates an expected call of AddResetToken.
func (mr *MockConnRunnerMockRecorder) AddResetToken(arg0, arg1 any) *ConnRunnerAddResetTokenCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddResetToken", reflect.TypeOf((*MockConnRunner)(nil).AddResetToken), arg0, arg1)
	return &ConnRunnerAddResetTokenCall{Call: call}
}

// ConnRunnerAddResetTokenCall wrap *gomock.Call
type ConnRunnerAddResetTokenCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ConnRunnerAddResetTokenCall) Return() *ConnRunnerAddResetTokenCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ConnRunnerAddResetTokenCall) Do(f func(protocol.StatelessResetToken, packetHandler)) *ConnRunnerAddResetTokenCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ConnRunnerAddResetTokenCall) DoAndReturn(f func(protocol.StatelessResetToken, packetHandler)) *ConnRunnerAddResetTokenCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetStatelessResetToken mocks base method.
func (m *MockConnRunner) GetStatelessResetToken(arg0 protocol.ConnectionID) protocol.StatelessResetToken {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatelessResetToken", arg0)
	ret0, _ := ret[0].(protocol.StatelessResetToken)
	return ret0
}

// GetStatelessResetToken indicates an expected call of GetStatelessResetToken.
func (mr *MockConnRunnerMockRecorder) GetStatelessResetToken(arg0 any) *ConnRunnerGetStatelessResetTokenCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatelessResetToken", reflect.TypeOf((*MockConnRunner)(nil).GetStatelessResetToken), arg0)
	return &ConnRunnerGetStatelessResetTokenCall{Call: call}
}

// ConnRunnerGetStatelessResetTokenCall wrap *gomock.Call
type ConnRunnerGetStatelessResetTokenCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ConnRunnerGetStatelessResetTokenCall) Return(arg0 protocol.StatelessResetToken) *ConnRunnerGetStatelessResetTokenCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ConnRunnerGetStatelessResetTokenCall) Do(f func(protocol.ConnectionID) protocol.StatelessResetToken) *ConnRunnerGetStatelessResetTokenCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ConnRunnerGetStatelessResetTokenCall) DoAndReturn(f func(protocol.ConnectionID) protocol.StatelessResetToken) *ConnRunnerGetStatelessResetTokenCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Remove mocks base method.
func (m *MockConnRunner) Remove(arg0 protocol.ConnectionID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Remove", arg0)
}

// Remove indicates an expected call of Remove.
func (mr *MockConnRunnerMockRecorder) Remove(arg0 any) *ConnRunnerRemoveCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockConnRunner)(nil).Remove), arg0)
	return &ConnRunnerRemoveCall{Call: call}
}

// ConnRunnerRemoveCall wrap *gomock.Call
type ConnRunnerRemoveCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ConnRunnerRemoveCall) Return() *ConnRunnerRemoveCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ConnRunnerRemoveCall) Do(f func(protocol.ConnectionID)) *ConnRunnerRemoveCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ConnRunnerRemoveCall) DoAndReturn(f func(protocol.ConnectionID)) *ConnRunnerRemoveCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RemoveResetToken mocks base method.
func (m *MockConnRunner) RemoveResetToken(arg0 protocol.StatelessResetToken) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveResetToken", arg0)
}

// RemoveResetToken indicates an expected call of RemoveResetToken.
func (mr *MockConnRunnerMockRecorder) RemoveResetToken(arg0 any) *ConnRunnerRemoveResetTokenCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveResetToken", reflect.TypeOf((*MockConnRunner)(nil).RemoveResetToken), arg0)
	return &ConnRunnerRemoveResetTokenCall{Call: call}
}

// ConnRunnerRemoveResetTokenCall wrap *gomock.Call
type ConnRunnerRemoveResetTokenCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ConnRunnerRemoveResetTokenCall) Return() *ConnRunnerRemoveResetTokenCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ConnRunnerRemoveResetTokenCall) Do(f func(protocol.StatelessResetToken)) *ConnRunnerRemoveResetTokenCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ConnRunnerRemoveResetTokenCall) DoAndReturn(f func(protocol.StatelessResetToken)) *ConnRunnerRemoveResetTokenCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ReplaceWithClosed mocks base method.
func (m *MockConnRunner) ReplaceWithClosed(arg0 []protocol.ConnectionID, arg1 protocol.Perspective, arg2 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReplaceWithClosed", arg0, arg1, arg2)
}

// ReplaceWithClosed indicates an expected call of ReplaceWithClosed.
func (mr *MockConnRunnerMockRecorder) ReplaceWithClosed(arg0, arg1, arg2 any) *ConnRunnerReplaceWithClosedCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplaceWithClosed", reflect.TypeOf((*MockConnRunner)(nil).ReplaceWithClosed), arg0, arg1, arg2)
	return &ConnRunnerReplaceWithClosedCall{Call: call}
}

// ConnRunnerReplaceWithClosedCall wrap *gomock.Call
type ConnRunnerReplaceWithClosedCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ConnRunnerReplaceWithClosedCall) Return() *ConnRunnerReplaceWithClosedCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ConnRunnerReplaceWithClosedCall) Do(f func([]protocol.ConnectionID, protocol.Perspective, []byte)) *ConnRunnerReplaceWithClosedCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ConnRunnerReplaceWithClosedCall) DoAndReturn(f func([]protocol.ConnectionID, protocol.Perspective, []byte)) *ConnRunnerReplaceWithClosedCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Retire mocks base method.
func (m *MockConnRunner) Retire(arg0 protocol.ConnectionID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Retire", arg0)
}

// Retire indicates an expected call of Retire.
func (mr *MockConnRunnerMockRecorder) Retire(arg0 any) *ConnRunnerRetireCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Retire", reflect.TypeOf((*MockConnRunner)(nil).Retire), arg0)
	return &ConnRunnerRetireCall{Call: call}
}

// ConnRunnerRetireCall wrap *gomock.Call
type ConnRunnerRetireCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ConnRunnerRetireCall) Return() *ConnRunnerRetireCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ConnRunnerRetireCall) Do(f func(protocol.ConnectionID)) *ConnRunnerRetireCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ConnRunnerRetireCall) DoAndReturn(f func(protocol.ConnectionID)) *ConnRunnerRetireCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
