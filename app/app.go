package app

import (
	"go-telemedicine/handler"
	"go-telemedicine/helpers"
	"go-telemedicine/repository"
	consultationrepository "go-telemedicine/repository/consultationRepository"
	schedulerepository "go-telemedicine/repository/scheduleRepository"
	userrepository "go-telemedicine/repository/userRepository"
	"go-telemedicine/service"
	consultationservice "go-telemedicine/service/consultationService"
	scheduleservice "go-telemedicine/service/scheduleService"
	userservice "go-telemedicine/service/userService"
)

func SetupApp(repo repository.Repository) handler.Handler {
	generator := helpers.NewGenerator()
	userRepo := userrepository.NewUserRepository(repo)
	userPermissionRepo := userrepository.NewUserPermissionRepository(repo)
	scheduleRepo := schedulerepository.NewScheduleRepository(repo)
	consultationRepo := consultationrepository.NewConsultationRepository(repo)

	service := service.NewService(generator, userRepo, userPermissionRepo, scheduleRepo, consultationRepo)

	userService := userservice.NewUserService(service)
	userPermissionService := userservice.NewUserPermissionService(service)
	scheduleService := scheduleservice.NewScheduleService(service)
	consultationService := consultationservice.NewConsultationService(service)

	handler := handler.NewHandler(userService, userPermissionService, scheduleService, consultationService)

	return handler
}
