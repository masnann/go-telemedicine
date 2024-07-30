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
	FindListUserRolePermissions(userID int64) ([]models.UserRolePermissionModels, error)
	AssignRoleToUserRequest(req models.AssignRoleToUserRequest) error
	FindUserRole(userID int64) (models.FindUserRoleResponse, error)
	CreatePermission(req models.PermissionModels) (int64, error)
	CreateRolePermission(req models.RolePermissionModels) (int64, error)
	CreateUserPermission(req models.UserPermissionModels) (int64, error)
	UserHavePermission(userID int64, permissionGroup, permissionName string) (bool, error)
	RoleHavePermission(userID int64, permissionGroup, permissionName string) (bool, error)
	FindPermissionsForUser(userID int64) ([]models.UserRolePermissionModels, error)
}

type ScheduleRepositoryInterface interface {
	CreateSchedule(req models.ScheduleModels) (int64, error)
	FindListAvailableSchedule(req models.ScheduleFindListAvailableRequest) ([]models.ScheduleModels, error)
	FindScheduleByID(id int64) (models.ScheduleModels, error)
}

type ConsultationRepositoryInterface interface {
	CreateConsultation(req models.ConsultationModels) (int64, error)
	FindListConsultationsUser(req models.ConsultationFindListByPatientIDRequest, userType string) ([]models.ConsultationModels, error)
}
