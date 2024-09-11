// Code generated by mockery v2.45.1. DO NOT EDIT.

package dao

import (
	context "context"

	api "github.com/content-services/content-sources-backend/pkg/api"

	mock "github.com/stretchr/testify/mock"
)

// MockRepositoryDao is an autogenerated mock type for the RepositoryDao type
type MockRepositoryDao struct {
	mock.Mock
}

// FetchForUrl provides a mock function with given fields: ctx, url
func (_m *MockRepositoryDao) FetchForUrl(ctx context.Context, url string) (Repository, error) {
	ret := _m.Called(ctx, url)

	if len(ret) == 0 {
		panic("no return value specified for FetchForUrl")
	}

	var r0 Repository
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (Repository, error)); ok {
		return rf(ctx, url)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) Repository); ok {
		r0 = rf(ctx, url)
	} else {
		r0 = ret.Get(0).(Repository)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchRepositoryRPMCount provides a mock function with given fields: ctx, repoUUID
func (_m *MockRepositoryDao) FetchRepositoryRPMCount(ctx context.Context, repoUUID string) (int, error) {
	ret := _m.Called(ctx, repoUUID)

	if len(ret) == 0 {
		panic("no return value specified for FetchRepositoryRPMCount")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (int, error)); ok {
		return rf(ctx, repoUUID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) int); ok {
		r0 = rf(ctx, repoUUID)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, repoUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListForIntrospection provides a mock function with given fields: ctx, urls, force
func (_m *MockRepositoryDao) ListForIntrospection(ctx context.Context, urls *[]string, force bool) ([]Repository, error) {
	ret := _m.Called(ctx, urls, force)

	if len(ret) == 0 {
		panic("no return value specified for ListForIntrospection")
	}

	var r0 []Repository
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *[]string, bool) ([]Repository, error)); ok {
		return rf(ctx, urls, force)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *[]string, bool) []Repository); ok {
		r0 = rf(ctx, urls, force)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Repository)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *[]string, bool) error); ok {
		r1 = rf(ctx, urls, force)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPublic provides a mock function with given fields: ctx, paginationData, _a2
func (_m *MockRepositoryDao) ListPublic(ctx context.Context, paginationData api.PaginationData, _a2 api.FilterData) (api.PublicRepositoryCollectionResponse, int64, error) {
	ret := _m.Called(ctx, paginationData, _a2)

	if len(ret) == 0 {
		panic("no return value specified for ListPublic")
	}

	var r0 api.PublicRepositoryCollectionResponse
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, api.PaginationData, api.FilterData) (api.PublicRepositoryCollectionResponse, int64, error)); ok {
		return rf(ctx, paginationData, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, api.PaginationData, api.FilterData) api.PublicRepositoryCollectionResponse); ok {
		r0 = rf(ctx, paginationData, _a2)
	} else {
		r0 = ret.Get(0).(api.PublicRepositoryCollectionResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, api.PaginationData, api.FilterData) int64); ok {
		r1 = rf(ctx, paginationData, _a2)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, api.PaginationData, api.FilterData) error); ok {
		r2 = rf(ctx, paginationData, _a2)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// OrphanCleanup provides a mock function with given fields: ctx
func (_m *MockRepositoryDao) OrphanCleanup(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for OrphanCleanup")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, repo
func (_m *MockRepositoryDao) Update(ctx context.Context, repo RepositoryUpdate) error {
	ret := _m.Called(ctx, repo)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, RepositoryUpdate) error); ok {
		r0 = rf(ctx, repo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockRepositoryDao creates a new instance of MockRepositoryDao. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRepositoryDao(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRepositoryDao {
	mock := &MockRepositoryDao{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
