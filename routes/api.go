package routes

import (
	"go-telemedicine/handler"
	consultationhandler "go-telemedicine/handler/consultationHandler"
	schedulehandler "go-telemedicine/handler/scheduleHandler"
	userhandler "go-telemedicine/handler/userHandler"
	"go-telemedicine/helpers/middleware"

	"github.com/labstack/echo/v4"
)

func ApiRoutes(e *echo.Echo, handler handler.Handler) {

	public := e.Group("/api/v1/public")
	userHandler := userhandler.NewUserHandler(handler)

	authGroup := public.Group("/auth")
	authGroup.POST("/register", userHandler.Register)
	authGroup.POST("/login", userHandler.Login)
	authGroup.POST("/token/refresh", userHandler.RefreshToken)

	private := e.Group("/api/v1/private")
	private.Use(middleware.JWTMiddleware)

	userGroup := private.Group("/user")
	userGroup.POST("/create", middleware.PermissionMiddleware(handler, "USER", "CREATE")(userHandler.CreateUser))
	userGroup.POST("/findbyid", middleware.PermissionMiddleware(handler, "USER", "READ")(userHandler.FindUserByID))
	userGroup.POST("/delete", middleware.PermissionMiddleware(handler, "USER", "DELETE")(userHandler.DeleteUser))
	userGroup.POST("/list", middleware.PermissionMiddleware(handler, "USER", "READ")(userHandler.FindListUsers))
	userGroup.POST("/rolepermission/create", middleware.PermissionMiddleware(handler, "ROLE_PERMISSION", "CREATE")(userHandler.CreateUserRolePermission))

	// Schedule
	scheduleHandler := schedulehandler.NewScheduleHandler(handler)
	scheduleGroup := private.Group("/schedule")
	scheduleGroup.POST("/create", middleware.PermissionMiddleware(handler, "SCHEDULE", "CREATE")(scheduleHandler.CreateSchedule))
	scheduleGroup.POST("/list", middleware.PermissionMiddleware(handler, "SCHEDULE", "READ")(scheduleHandler.FindListAvailableSchedule))

	// Consultation
	consultationHandler := consultationhandler.NewConsultationHandler(handler)
	consultationGroup := private.Group("/consultation")
	consultationGroup.POST("/create", middleware.PermissionMiddleware(handler, "CONSULTATION", "CREATE")(consultationHandler.CreateConsultation))
	consultationGroup.POST("/list/bypatientid", consultationHandler.FindListConsultationsByPatientID)
	consultationGroup.POST("/list/bydoctorid", consultationHandler.FindListConsultationsByDoctorID)
}
