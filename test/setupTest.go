package test

import (
	helpers "go-telemedicine/helpers/mocks"
	"go-telemedicine/repository/mocks"
	"go-telemedicine/service"
	consultationservice "go-telemedicine/service/consultationService"
	scheduleservice "go-telemedicine/service/scheduleService"
	userservice "go-telemedicine/service/userService"
	"testing"
)

type TestSuite struct {
	Generator          *helpers.GeneratorInterface
	UserRepo           *mocks.UserRepositoryInterface
	UserPermissionRepo *mocks.UserPermissionRepositoryInterface
	ScheduleRepo       *mocks.ScheduleRepositoryInterface
	ConsultationRepo   *mocks.ConsultationRepositoryInterface

	Service               service.Service
	UserService           userservice.UserService
	UserPermissionService userservice.UserPermissionService
	ScheduleService       scheduleservice.ScheduleService
	ConsultationService   consultationservice.ConsultationService
}

func SetupTestCase(t *testing.T) *TestSuite {

	generator := helpers.NewGeneratorInterface(t)
	userRepo := mocks.NewUserRepositoryInterface(t)
	userPermissionRepo := mocks.NewUserPermissionRepositoryInterface(t)
	scheduleRepo := mocks.NewScheduleRepositoryInterface(t)
	consultationRepo := mocks.NewConsultationRepositoryInterface(t)

	svc := service.NewService(generator, userRepo, userPermissionRepo, scheduleRepo, consultationRepo)

	userService := userservice.NewUserService(svc)
	userPermissionService := userservice.NewUserPermissionService(svc)
	scheduleService := scheduleservice.NewScheduleService(svc)
	consultationService := consultationservice.NewConsultationService(svc)

	return &TestSuite{
		Generator:          generator,
		UserRepo:           userRepo,
		UserPermissionRepo: userPermissionRepo,
		ScheduleRepo:       scheduleRepo,
		ConsultationRepo:   consultationRepo,

		Service:               svc,
		UserService:           userService,
		UserPermissionService: userPermissionService,
		ScheduleService:       scheduleService,
		ConsultationService:   consultationService,
	}
}
