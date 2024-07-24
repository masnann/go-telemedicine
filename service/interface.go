package service

import (
	"go-telemedicine/models"
)

type UserServiceInterface interface {
	Register(req models.UserRegisterRequest) (int64, error)
	FindUserByID(req models.RequestID) (models.UserModels, error)
	Login(req models.UserLoginRequest) (models.UserLoginResponse, error)
	RefreshToken(accessToken string) (models.UserLoginResponse, error)
}

type UserPermissionServiceInterface interface {
	FindListUserPermissions(userID int64) ([]models.UserPermissionModels, error)
	FindUserPermissions(userID int64, permissionGroup, permissionName string) (models.UserPermissionModels, error)
}
