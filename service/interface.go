package service

import (
	"go-telemedicine/models"
)

type UserServiceInterface interface {
	Register(req models.UserRegisterRequest) (int64, error)
	FindUserByID(req models.RequestID) (models.UserModels, error)
	Login(req models.UserLoginRequest) (models.UserLoginResponse, error)
	RefreshToken(accessToken string) (models.UserLoginResponse, error)
	DeleteUser(req models.RequestID) error
	FindListUsers(req models.FindListUserRequest) ([]models.FindListUserResponse, error)
	CreateUser(req models.UserCreateRequest) (int64, error)
}

type UserPermissionServiceInterface interface {
	FindListUserPermissions(userID int64) ([]models.UserPermissionModels, error)
	FindUserPermissions(userID int64, permissionGroup, permissionName string) (models.UserPermissionModels, error)
	CreateUserRolePermission(req models.UserRolePermissionCreateRequest) (int64, error) 
}

type ScheduleServiceInterface interface {
	CreateSchedule(req models.ScheduleCreateRequest) (int64, error)
	FindListAvailableSchedule(req models.ScheduleFindListAvailableRequest) ([]models.ScheduleModels, error)
}

type ConsultationServiceInterface interface {
    CreateConsultation(req models.ConsultationCreateRequest) (int64, error)
}