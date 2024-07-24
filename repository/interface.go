package repository

import "go-telemedicine/models"

type UserRepositoryInterface interface {
	Register(req models.UserModels) (int64, error)
	FindUserByID(id int64) (models.UserModels, error)
	Login(email string) (models.UserModels, error)
}

type UserPermissionRepositoryInterface interface {
	FindListUserPermissions(userID int64) ([]models.UserPermissionModels, error)
	AssignRoleToUserRequest(req models.AssignRoleToUserRequest) error
	FindUserRole(userID int64) (models.FindUserRoleResponse, error)
	FindUserPermissions(userID int64, permissionGroup, permissionName string) (models.UserPermissionModels, error)
}
