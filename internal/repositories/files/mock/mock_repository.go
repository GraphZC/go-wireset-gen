// Code generated by mockery. DO NOT EDIT.

package mock_files

import mock "github.com/stretchr/testify/mock"

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

type Repository_Expecter struct {
	mock *mock.Mock
}

func (_m *Repository) EXPECT() *Repository_Expecter {
	return &Repository_Expecter{mock: &_m.Mock}
}

// GetGoModFile provides a mock function with no fields
func (_m *Repository) GetGoModFile() (string, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetGoModFile")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_GetGoModFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetGoModFile'
type Repository_GetGoModFile_Call struct {
	*mock.Call
}

// GetGoModFile is a helper method to define mock.On call
func (_e *Repository_Expecter) GetGoModFile() *Repository_GetGoModFile_Call {
	return &Repository_GetGoModFile_Call{Call: _e.mock.On("GetGoModFile")}
}

func (_c *Repository_GetGoModFile_Call) Run(run func()) *Repository_GetGoModFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Repository_GetGoModFile_Call) Return(_a0 string, _a1 error) *Repository_GetGoModFile_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_GetGoModFile_Call) RunAndReturn(run func() (string, error)) *Repository_GetGoModFile_Call {
	_c.Call.Return(run)
	return _c
}

// ListAllGoFiles provides a mock function with no fields
func (_m *Repository) ListAllGoFiles() ([]string, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ListAllGoFiles")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_ListAllGoFiles_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAllGoFiles'
type Repository_ListAllGoFiles_Call struct {
	*mock.Call
}

// ListAllGoFiles is a helper method to define mock.On call
func (_e *Repository_Expecter) ListAllGoFiles() *Repository_ListAllGoFiles_Call {
	return &Repository_ListAllGoFiles_Call{Call: _e.mock.On("ListAllGoFiles")}
}

func (_c *Repository_ListAllGoFiles_Call) Run(run func()) *Repository_ListAllGoFiles_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Repository_ListAllGoFiles_Call) Return(_a0 []string, _a1 error) *Repository_ListAllGoFiles_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_ListAllGoFiles_Call) RunAndReturn(run func() ([]string, error)) *Repository_ListAllGoFiles_Call {
	_c.Call.Return(run)
	return _c
}

// ReadFile provides a mock function with given fields: filePath
func (_m *Repository) ReadFile(filePath string) (string, error) {
	ret := _m.Called(filePath)

	if len(ret) == 0 {
		panic("no return value specified for ReadFile")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(filePath)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(filePath)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(filePath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_ReadFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadFile'
type Repository_ReadFile_Call struct {
	*mock.Call
}

// ReadFile is a helper method to define mock.On call
//   - filePath string
func (_e *Repository_Expecter) ReadFile(filePath interface{}) *Repository_ReadFile_Call {
	return &Repository_ReadFile_Call{Call: _e.mock.On("ReadFile", filePath)}
}

func (_c *Repository_ReadFile_Call) Run(run func(filePath string)) *Repository_ReadFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Repository_ReadFile_Call) Return(_a0 string, _a1 error) *Repository_ReadFile_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_ReadFile_Call) RunAndReturn(run func(string) (string, error)) *Repository_ReadFile_Call {
	_c.Call.Return(run)
	return _c
}

// WriteFile provides a mock function with given fields: directory, fileName, data
func (_m *Repository) WriteFile(directory string, fileName string, data string) error {
	ret := _m.Called(directory, fileName, data)

	if len(ret) == 0 {
		panic("no return value specified for WriteFile")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(directory, fileName, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_WriteFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteFile'
type Repository_WriteFile_Call struct {
	*mock.Call
}

// WriteFile is a helper method to define mock.On call
//   - directory string
//   - fileName string
//   - data string
func (_e *Repository_Expecter) WriteFile(directory interface{}, fileName interface{}, data interface{}) *Repository_WriteFile_Call {
	return &Repository_WriteFile_Call{Call: _e.mock.On("WriteFile", directory, fileName, data)}
}

func (_c *Repository_WriteFile_Call) Run(run func(directory string, fileName string, data string)) *Repository_WriteFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *Repository_WriteFile_Call) Return(_a0 error) *Repository_WriteFile_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_WriteFile_Call) RunAndReturn(run func(string, string, string) error) *Repository_WriteFile_Call {
	_c.Call.Return(run)
	return _c
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
