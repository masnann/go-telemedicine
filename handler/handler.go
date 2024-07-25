package handler

import "go-telemedicine/service"

type Handler struct {
	UserService           service.UserServiceInterface
	UserPermissionService service.UserPermissionServiceInterface
	ScheduleService       service.ScheduleServiceInterface
	ConsultationService   service.ConsultationServiceInterface
}

func NewHandler(
	userService service.UserServiceInterface,
	userPermissionService service.UserPermissionServiceInterface,
	scheduleService service.ScheduleServiceInterface,
	consultationService service.ConsultationServiceInterface,

) Handler {
	return Handler{
		UserService:           userService,
		UserPermissionService: userPermissionService,
		ScheduleService:       scheduleService,
		ConsultationService:   consultationService,
	}
}
