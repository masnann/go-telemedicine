package middleware

import (
	"go-telemedicine/config"
	"go-telemedicine/constants"
	"go-telemedicine/handler"
	"go-telemedicine/helpers"
	"go-telemedicine/models"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// JWTMiddleware validates the JWT token and sets the user in the context
func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var result models.Response
		// Extract the JWT token from the request header
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			result = helpers.ResponseJSON(false, constants.VALIDATE_ERROR_CODE, "Missing authorization header", nil)
			return c.JSON(http.StatusBadRequest, result)
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JWTSecret), nil
		})
		if err != nil {
			result = helpers.ResponseJSON(false, constants.UNAUTHORIZED_CODE, err.Error(), nil)
			return c.JSON(http.StatusUnauthorized, result)
		}

		// Extract claims and set user role in the context
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := claims["user_id"].(float64)
			c.Set("user_id", int64(userID))

			userRole := claims["role"].(string)
			c.Set("role", userRole)

		} else {
			result = helpers.ResponseJSON(false, constants.UNAUTHORIZED_CODE, "Invalid token", nil)
			return c.JSON(http.StatusUnauthorized, result)
		}

		return next(c)
	}
}

// Permission Middleware to check user permissions
func PermissionMiddleware(handler handler.Handler, permissionGroup, permissionName string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var result models.Response
			userID := c.Get("user_id").(int64)

			permission, err := handler.UserPermissionService.FindUserPermissions(userID, permissionGroup, permissionName)
			if err != nil {
				log.Printf("Error PermissionMiddleware: %v", err)
				result = helpers.ResponseJSON(false, constants.SYSTEM_ERROR_CODE, err.Error(), nil)
				return c.JSON(http.StatusInternalServerError, result)
			}
			if !permission.Status {
				log.Printf("Permission status: %v", permission.Status)
				result = helpers.ResponseJSON(false, constants.UNAUTHORIZED_CODE, "Access denied. You don't have permission", nil)
				return c.JSON(http.StatusUnauthorized, result)
			}

			return next(c)
		}
	}
}
