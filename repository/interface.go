package repository

import "go-telemedicine/models"

type UserRepositoryInterface interface {
	Register(req models.UserModels) (int64, error)
	FindUserByID(id int64) (models.UserModels, error)
	Login(email string) (models.UserModels, error)
	DeleteUser(userID int64) error
	FindListUser(req models.FindListUserRequest) ([]models.FindListUserResponse, error)
}

type UserPermissionRepositoryInterface interface {
	FindListUserPermissions(userID int64) ([]models.UserPermissionModels, error)
	AssignRoleToUserRequest(req models.AssignRoleToUserRequest) error
	FindUserRole(userID int64) (models.FindUserRoleResponse, error)
	FindUserPermissions(userID int64, permissionGroup, permissionName string) (models.UserPermissionModels, error)
	CreateUserRolePermission(req models.RolePermissionModels) (int64, error)
}

type ScheduleRepositoryInterface interface {
	CreateSchedule(req models.ScheduleModels) (int64, error)
	FindListAvailableSchedule(req models.ScheduleFindListAvailableRequest) ([]models.ScheduleModels, error)
}
