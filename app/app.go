package app

import (
	"go-telemedicine/handler"
	"go-telemedicine/helpers"
	"go-telemedicine/repository"
	schedulerepository "go-telemedicine/repository/scheduleRepository"
	userrepository "go-telemedicine/repository/userRepository"
	"go-telemedicine/service"
	scheduleservice "go-telemedicine/service/scheduleService"
	userservice "go-telemedicine/service/userService"
)

func SetupApp(repo repository.Repository) handler.Handler {
	generator := helpers.NewGenerator()
	userRepo := userrepository.NewUserRepository(repo)
	userPermissionRepo := userrepository.NewUserPermissionRepository(repo)
	scheduleRepo := schedulerepository.NewScheduleRepository(repo)

	service := service.NewService(generator, userRepo, userPermissionRepo, scheduleRepo)

	userService := userservice.NewUserService(service)
	userPermissionService := userservice.NewUserPermissionService(service)
	scheduleService := scheduleservice.NewScheduleService(service)

	handler := handler.NewHandler(userService, userPermissionService, scheduleService)

	return handler
}
