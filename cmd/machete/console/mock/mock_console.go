// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/docker/libmachete/cmd/machete/console (interfaces: Console)

package mock

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of Console interface
type MockConsole struct {
	ctrl     *gomock.Controller
	recorder *_MockConsoleRecorder
}

// Recorder for MockConsole (not exported)
type _MockConsoleRecorder struct {
	mock *MockConsole
}

func NewMockConsole(ctrl *gomock.Controller) *MockConsole {
	mock := &MockConsole{ctrl: ctrl}
	mock.recorder = &_MockConsoleRecorder{mock}
	return mock
}

func (_m *MockConsole) EXPECT() *_MockConsoleRecorder {
	return _m.recorder
}

func (_m *MockConsole) ErrPrintln(_param0 ...interface{}) {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "ErrPrintln", _s...)
}

func (_mr *_MockConsoleRecorder) ErrPrintln(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ErrPrintln", arg0...)
}

func (_m *MockConsole) Fatal(_param0 ...interface{}) {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Fatal", _s...)
}

func (_mr *_MockConsoleRecorder) Fatal(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Fatal", arg0...)
}

func (_m *MockConsole) Println(_param0 ...interface{}) {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Println", _s...)
}

func (_mr *_MockConsoleRecorder) Println(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Println", arg0...)
}