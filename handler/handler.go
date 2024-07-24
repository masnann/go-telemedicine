package handler

import "go-telemedicine/service"

type Handler struct {
	UserService           service.UserServiceInterface
	UserPermissionService service.UserPermissionServiceInterface
	ScheduleService       service.ScheduleServiceInterface
}

func NewHandler(
	userService service.UserServiceInterface,
	userPermissionService service.UserPermissionServiceInterface,
	scheduleService service.ScheduleServiceInterface,

) Handler {
	return Handler{
		UserService:           userService,
		UserPermissionService: userPermissionService,
		ScheduleService:       scheduleService,
	}
}
