package service

import (
	"go-telemedicine/helpers"
	"go-telemedicine/repository"
)

type Service struct {
	Generator          helpers.GeneratorInterface
	UserRepo           repository.UserRepositoryInterface
	UserPermissionRepo repository.UserPermissionRepositoryInterface
	ScheduleRepo       repository.ScheduleRepositoryInterface
}

func NewService(
	generator helpers.GeneratorInterface,
	userRepo repository.UserRepositoryInterface,
	userPermissionRepo repository.UserPermissionRepositoryInterface,
	scheduleRepo repository.ScheduleRepositoryInterface,
) Service {
	return Service{
		Generator:          generator,
		UserRepo:           userRepo,
		UserPermissionRepo: userPermissionRepo,
		ScheduleRepo:       scheduleRepo,
	}
}
