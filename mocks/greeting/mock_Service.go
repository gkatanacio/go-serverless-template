// Code generated by mockery. DO NOT EDIT.

package greeting

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockService is an autogenerated mock type for the Service type
type MockService struct {
	mock.Mock
}

type MockService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockService) EXPECT() *MockService_Expecter {
	return &MockService_Expecter{mock: &_m.Mock}
}

// HelloMessage provides a mock function with given fields: ctx, name
func (_m *MockService) HelloMessage(ctx context.Context, name string) string {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for HelloMessage")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockService_HelloMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HelloMessage'
type MockService_HelloMessage_Call struct {
	*mock.Call
}

// HelloMessage is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockService_Expecter) HelloMessage(ctx interface{}, name interface{}) *MockService_HelloMessage_Call {
	return &MockService_HelloMessage_Call{Call: _e.mock.On("HelloMessage", ctx, name)}
}

func (_c *MockService_HelloMessage_Call) Run(run func(ctx context.Context, name string)) *MockService_HelloMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockService_HelloMessage_Call) Return(_a0 string) *MockService_HelloMessage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockService_HelloMessage_Call) RunAndReturn(run func(context.Context, string) string) *MockService_HelloMessage_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockService creates a new instance of MockService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockService {
	mock := &MockService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
