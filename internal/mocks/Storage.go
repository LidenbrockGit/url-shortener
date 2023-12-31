// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	linkentity "github.com/LidenbrockGit/url-shortener/internal/entities/linkentity"
	uuid "github.com/google/uuid"
	mock "github.com/stretchr/testify/mock"
)

// Storage is an autogenerated mock type for the Storage type
type Storage struct {
	mock.Mock
}

type Storage_Expecter struct {
	mock *mock.Mock
}

func (_m *Storage) EXPECT() *Storage_Expecter {
	return &Storage_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, link
func (_m *Storage) Create(ctx context.Context, link linkentity.Link) (linkentity.Link, error) {
	ret := _m.Called(ctx, link)

	var r0 linkentity.Link
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, linkentity.Link) (linkentity.Link, error)); ok {
		return rf(ctx, link)
	}
	if rf, ok := ret.Get(0).(func(context.Context, linkentity.Link) linkentity.Link); ok {
		r0 = rf(ctx, link)
	} else {
		r0 = ret.Get(0).(linkentity.Link)
	}

	if rf, ok := ret.Get(1).(func(context.Context, linkentity.Link) error); ok {
		r1 = rf(ctx, link)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Storage_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type Storage_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - link linkentity.Link
func (_e *Storage_Expecter) Create(ctx interface{}, link interface{}) *Storage_Create_Call {
	return &Storage_Create_Call{Call: _e.mock.On("Create", ctx, link)}
}

func (_c *Storage_Create_Call) Run(run func(ctx context.Context, link linkentity.Link)) *Storage_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(linkentity.Link))
	})
	return _c
}

func (_c *Storage_Create_Call) Return(_a0 linkentity.Link, _a1 error) *Storage_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Storage_Create_Call) RunAndReturn(run func(context.Context, linkentity.Link) (linkentity.Link, error)) *Storage_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, linkId
func (_m *Storage) Delete(ctx context.Context, linkId uuid.UUID) (linkentity.Link, error) {
	ret := _m.Called(ctx, linkId)

	var r0 linkentity.Link
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (linkentity.Link, error)); ok {
		return rf(ctx, linkId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) linkentity.Link); ok {
		r0 = rf(ctx, linkId)
	} else {
		r0 = ret.Get(0).(linkentity.Link)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, linkId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Storage_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type Storage_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - linkId uuid.UUID
func (_e *Storage_Expecter) Delete(ctx interface{}, linkId interface{}) *Storage_Delete_Call {
	return &Storage_Delete_Call{Call: _e.mock.On("Delete", ctx, linkId)}
}

func (_c *Storage_Delete_Call) Run(run func(ctx context.Context, linkId uuid.UUID)) *Storage_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *Storage_Delete_Call) Return(_a0 linkentity.Link, _a1 error) *Storage_Delete_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Storage_Delete_Call) RunAndReturn(run func(context.Context, uuid.UUID) (linkentity.Link, error)) *Storage_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Read provides a mock function with given fields: ctx, linkId
func (_m *Storage) Read(ctx context.Context, linkId uuid.UUID) (linkentity.Link, error) {
	ret := _m.Called(ctx, linkId)

	var r0 linkentity.Link
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (linkentity.Link, error)); ok {
		return rf(ctx, linkId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) linkentity.Link); ok {
		r0 = rf(ctx, linkId)
	} else {
		r0 = ret.Get(0).(linkentity.Link)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, linkId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Storage_Read_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Read'
type Storage_Read_Call struct {
	*mock.Call
}

// Read is a helper method to define mock.On call
//   - ctx context.Context
//   - linkId uuid.UUID
func (_e *Storage_Expecter) Read(ctx interface{}, linkId interface{}) *Storage_Read_Call {
	return &Storage_Read_Call{Call: _e.mock.On("Read", ctx, linkId)}
}

func (_c *Storage_Read_Call) Run(run func(ctx context.Context, linkId uuid.UUID)) *Storage_Read_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *Storage_Read_Call) Return(_a0 linkentity.Link, _a1 error) *Storage_Read_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Storage_Read_Call) RunAndReturn(run func(context.Context, uuid.UUID) (linkentity.Link, error)) *Storage_Read_Call {
	_c.Call.Return(run)
	return _c
}

// ReadAll provides a mock function with given fields: ctx
func (_m *Storage) ReadAll(ctx context.Context) (<-chan linkentity.Link, error) {
	ret := _m.Called(ctx)

	var r0 <-chan linkentity.Link
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (<-chan linkentity.Link, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) <-chan linkentity.Link); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan linkentity.Link)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Storage_ReadAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadAll'
type Storage_ReadAll_Call struct {
	*mock.Call
}

// ReadAll is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Storage_Expecter) ReadAll(ctx interface{}) *Storage_ReadAll_Call {
	return &Storage_ReadAll_Call{Call: _e.mock.On("ReadAll", ctx)}
}

func (_c *Storage_ReadAll_Call) Run(run func(ctx context.Context)) *Storage_ReadAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Storage_ReadAll_Call) Return(_a0 <-chan linkentity.Link, _a1 error) *Storage_ReadAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Storage_ReadAll_Call) RunAndReturn(run func(context.Context) (<-chan linkentity.Link, error)) *Storage_ReadAll_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, link
func (_m *Storage) Update(ctx context.Context, link linkentity.Link) error {
	ret := _m.Called(ctx, link)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, linkentity.Link) error); ok {
		r0 = rf(ctx, link)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Storage_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type Storage_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - link linkentity.Link
func (_e *Storage_Expecter) Update(ctx interface{}, link interface{}) *Storage_Update_Call {
	return &Storage_Update_Call{Call: _e.mock.On("Update", ctx, link)}
}

func (_c *Storage_Update_Call) Run(run func(ctx context.Context, link linkentity.Link)) *Storage_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(linkentity.Link))
	})
	return _c
}

func (_c *Storage_Update_Call) Return(_a0 error) *Storage_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Storage_Update_Call) RunAndReturn(run func(context.Context, linkentity.Link) error) *Storage_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewStorage creates a new instance of Storage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *Storage {
	mock := &Storage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
