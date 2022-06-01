// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	entity "bareksa-take-home-test-michael-koh/core/entity"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// NewsService is an autogenerated mock type for the NewsService type
type NewsService struct {
	mock.Mock
}

// CreateNews provides a mock function with given fields: ctx, newNews
func (_m *NewsService) CreateNews(ctx context.Context, newNews entity.News) error {
	ret := _m.Called(ctx, newNews)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.News) error); ok {
		r0 = rf(ctx, newNews)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteNews provides a mock function with given fields: ctx, newsSerial
func (_m *NewsService) DeleteNews(ctx context.Context, newsSerial string) error {
	ret := _m.Called(ctx, newsSerial)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, newsSerial)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetNews provides a mock function with given fields: ctx, newsQuery
func (_m *NewsService) GetNews(ctx context.Context, newsQuery entity.GetNewsQuery) ([]entity.News, error) {
	ret := _m.Called(ctx, newsQuery)

	var r0 []entity.News
	if rf, ok := ret.Get(0).(func(context.Context, entity.GetNewsQuery) []entity.News); ok {
		r0 = rf(ctx, newsQuery)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.News)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.GetNewsQuery) error); ok {
		r1 = rf(ctx, newsQuery)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateNews provides a mock function with given fields: ctx, newNews
func (_m *NewsService) UpdateNews(ctx context.Context, newNews entity.News) error {
	ret := _m.Called(ctx, newNews)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.News) error); ok {
		r0 = rf(ctx, newNews)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type NewNewsServiceT interface {
	mock.TestingT
	Cleanup(func())
}

// NewNewsService creates a new instance of NewsService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNewsService(t NewNewsServiceT) *NewsService {
	mock := &NewsService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}