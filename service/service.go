package service

import (
	"go-telemedicine/helpers"
	"go-telemedicine/repository"
)

type Service struct {
	Generator          helpers.GeneratorInterface
	UserRepo           repository.UserRepositoryInterface
	UserPermissionRepo repository.UserPermissionRepositoryInterface
}

func NewService(
	generator helpers.GeneratorInterface,
	userRepo repository.UserRepositoryInterface,
	userPermissionRepo repository.UserPermissionRepositoryInterface,
) Service {
	return Service{
		Generator:          generator,
		UserRepo:           userRepo,
		UserPermissionRepo: userPermissionRepo,
	}
}
