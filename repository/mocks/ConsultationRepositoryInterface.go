// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	models "go-telemedicine/models"

	mock "github.com/stretchr/testify/mock"
)

// ConsultationRepositoryInterface is an autogenerated mock type for the ConsultationRepositoryInterface type
type ConsultationRepositoryInterface struct {
	mock.Mock
}

// CreateConsultation provides a mock function with given fields: req
func (_m *ConsultationRepositoryInterface) CreateConsultation(req models.ConsultationModels) (int64, error) {
	ret := _m.Called(req)

	if len(ret) == 0 {
		panic("no return value specified for CreateConsultation")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(models.ConsultationModels) (int64, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(models.ConsultationModels) int64); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(models.ConsultationModels) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindListConsultationsUser provides a mock function with given fields: req, userType
func (_m *ConsultationRepositoryInterface) FindListConsultationsUser(req models.ConsultationFindListByPatientIDRequest, userType string) ([]models.ConsultationModels, error) {
	ret := _m.Called(req, userType)

	if len(ret) == 0 {
		panic("no return value specified for FindListConsultationsUser")
	}

	var r0 []models.ConsultationModels
	var r1 error
	if rf, ok := ret.Get(0).(func(models.ConsultationFindListByPatientIDRequest, string) ([]models.ConsultationModels, error)); ok {
		return rf(req, userType)
	}
	if rf, ok := ret.Get(0).(func(models.ConsultationFindListByPatientIDRequest, string) []models.ConsultationModels); ok {
		r0 = rf(req, userType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.ConsultationModels)
		}
	}

	if rf, ok := ret.Get(1).(func(models.ConsultationFindListByPatientIDRequest, string) error); ok {
		r1 = rf(req, userType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewConsultationRepositoryInterface creates a new instance of ConsultationRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConsultationRepositoryInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ConsultationRepositoryInterface {
	mock := &ConsultationRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
