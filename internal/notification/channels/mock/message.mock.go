// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/zitadel/zitadel/internal/notification/channels (interfaces: Message)
//
// Generated by this command:
//
//	mockgen -package mock -destination ./mock/message.mock.go github.com/zitadel/zitadel/internal/notification/channels Message
//

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	eventstore "github.com/zitadel/zitadel/internal/eventstore"
	gomock "go.uber.org/mock/gomock"
)

// MockMessage is a mock of Message interface.
type MockMessage struct {
	ctrl     *gomock.Controller
	recorder *MockMessageMockRecorder
}

// MockMessageMockRecorder is the mock recorder for MockMessage.
type MockMessageMockRecorder struct {
	mock *MockMessage
}

// NewMockMessage creates a new mock instance.
func NewMockMessage(ctrl *gomock.Controller) *MockMessage {
	mock := &MockMessage{ctrl: ctrl}
	mock.recorder = &MockMessageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessage) EXPECT() *MockMessageMockRecorder {
	return m.recorder
}

// GetContent mocks base method.
func (m *MockMessage) GetContent() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContent")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContent indicates an expected call of GetContent.
func (mr *MockMessageMockRecorder) GetContent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContent", reflect.TypeOf((*MockMessage)(nil).GetContent))
}

// GetTriggeringEvent mocks base method.
func (m *MockMessage) GetTriggeringEvent() eventstore.Event {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTriggeringEvent")
	ret0, _ := ret[0].(eventstore.Event)
	return ret0
}

// GetTriggeringEvent indicates an expected call of GetTriggeringEvent.
func (mr *MockMessageMockRecorder) GetTriggeringEvent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTriggeringEvent", reflect.TypeOf((*MockMessage)(nil).GetTriggeringEvent))
}
