package routes

import (
	"go-telemedicine/handler"
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
	userGroup.POST("/findbyid", middleware.PermissionMiddleware(handler, "USER", "READ")(userHandler.FindUserByID))
	userGroup.POST("/delete", middleware.PermissionMiddleware(handler, "USER", "DELETE")(userHandler.DeleteUser))
	userGroup.POST("/list", middleware.PermissionMiddleware(handler, "USER", "READ")(userHandler.FindListUsers))
}
