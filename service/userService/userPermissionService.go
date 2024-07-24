package userservice

import (
	"errors"
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

func (s UserPermissionService) FindListUserPermissions(userID int64) ([]models.UserPermissionModels, error) {
	result, err := s.service.UserPermissionRepo.FindListUserPermissions(userID)
	if err != nil {
		log.Println("Error finding user permissions: ", err)
		return nil, errors.New("failed to find user permissions")
	}
	return result, nil
}

func (s UserPermissionService) FindUserPermissions(userID int64, permissionGroup, permissionName string) (models.UserPermissionModels, error) {
	result, err := s.service.UserPermissionRepo.FindUserPermissions(userID, permissionGroup, permissionName)
	if err != nil {
		log.Println("Error finding user permissions: ", err)
		return result, errors.New("failed to find user permissions")
	}
	return result, nil

}
