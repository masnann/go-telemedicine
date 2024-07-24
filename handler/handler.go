package handler

import "go-telemedicine/service"

type Handler struct {
	UserService           service.UserServiceInterface
	UserPermissionService service.UserPermissionServiceInterface
}

func NewHandler(
	userService service.UserServiceInterface,
	userPermissionService service.UserPermissionServiceInterface,
) Handler {
	return Handler{
		UserService:           userService,
		UserPermissionService: userPermissionService,
	}
}
