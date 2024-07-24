package app

import (
	"go-telemedicine/handler"
	"go-telemedicine/helpers"
	"go-telemedicine/repository"
	userrepository "go-telemedicine/repository/userRepository"
	"go-telemedicine/service"
	userservice "go-telemedicine/service/userService"
)

func SetupApp(repo repository.Repository) handler.Handler {
	generator := helpers.NewGenerator()
	userRepo := userrepository.NewUserRepository(repo)
	userPermissionRepo := userrepository.NewUserPermissionRepository(repo)

	service := service.NewService(generator, userRepo, userPermissionRepo)

	userService := userservice.NewUserService(service)
	userPermissionService := userservice.NewUserPermissionService(service)

	handler := handler.NewHandler(userService, userPermissionService)

	return handler
}
