package userservice

import (
	"errors"
	"go-telemedicine/helpers"
	"go-telemedicine/models"
	"go-telemedicine/service"
	"log"
)

type UserPermissionService struct {
	service service.Service
}

func NewUserPermissionService(service service.Service) UserPermissionService {
	return UserPermissionService{
		service: service,
	}
}

func (s UserPermissionService) FindListUserRolePermissions(userID int64) ([]models.UserRolePermissionModels, error) {
	result, err := s.service.UserPermissionRepo.FindListUserRolePermissions(userID)
	if err != nil {
		log.Println("Error finding user permissions: ", err)
		return nil, errors.New("failed to find user permissions")
	}
	return result, nil
}

func (s UserPermissionService) CreatePermission(req models.PermissionCreateRequest) (int64, error) {
	newData := models.PermissionModels{
		Group:     req.Group,
		Name:      req.Name,
		CreatedAt: helpers.TimeStampNow(),
		UpdatedAt: "",
	}
	result, err := s.service.UserPermissionRepo.CreatePermission(newData)
	if err != nil {
		log.Println("Error creating user role permission: ", err)
		return 0, err
	}
	return result, nil
}

func (s UserPermissionService) CreateRolePermission(req models.RolePermissionCreateRequest) (int64, error) {
	newData := models.RolePermissionModels{
		RoleID:       req.RoleID,
		PermissionID: req.PermissionID,
	}
	result, err := s.service.UserPermissionRepo.CreateRolePermission(newData)
	if err != nil {
		log.Println("Error creating role permission: ", err)
		return 0, err
	}
	return result, nil

}

func (s UserPermissionService) CreateUserPermission(req models.UserPermissionCreateRequest) (int64, error) {
	admin, err := s.service.UserRepo.FindUserByID(req.AdminID)
	if err != nil {
		log.Println("Error finding admin by ID: ", err)
		return 0, errors.New("admin not found")
	}
	newData := models.UserPermissionModels{
		UserID:       req.UserID,
		PermissionID: req.PermissionID,
		Status:       req.Status,
		GrantedBy:    admin.ID,
		GrantedAt:    helpers.TimeStampNow(),
		UpdatedAt:    "",
	}

	result, err := s.service.UserPermissionRepo.CreateUserPermission(newData)
	if err != nil {
		log.Println("Error creating user permission: ", err)
		return 0, err
	}
	return result, nil

}

func (s UserPermissionService) UserHavePermission(userID int64, permissionGroup, permissionName string) (bool, error) {
	result, err := s.service.UserPermissionRepo.UserHavePermission(userID, permissionGroup, permissionName)
	if err != nil {
		log.Println("Error checking permission: ", err)
		return false, err
	}
	return result, nil
}

func (s UserPermissionService) RoleHavePermission(userID int64, permissionGroup, permissionName string) (bool, error) {
	result, err := s.service.UserPermissionRepo.RoleHavePermission(userID, permissionGroup, permissionName)
	if err != nil {
		log.Println("Error checking role permission: ", err)
		return false, err
	}
	return result, nil
}
