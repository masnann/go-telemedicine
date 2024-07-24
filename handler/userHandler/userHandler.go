package userhandler

import (
	"go-telemedicine/constants"
	"go-telemedicine/handler"
	"go-telemedicine/helpers"
	"go-telemedicine/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	handler handler.Handler
}

func NewUserHandler(handler handler.Handler) UserHandler {
	return UserHandler{
		handler: handler,
	}
}

func (h UserHandler) Register(ctx echo.Context) error {
	var result models.Response

	req := new(models.UserRegisterRequest)
	if err := helpers.ValidateStruct(ctx, req); err != nil {
		log.Printf("Error Failed to validate request: %v", err)
		result = helpers.ResponseJSON(false, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	userID, err := h.handler.UserService.Register(*req)
	if err != nil {
		log.Printf("Error Register: %v", err)
		result = helpers.ResponseJSON(false, constants.SYSTEM_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusInternalServerError, result)
	}
	result = helpers.ResponseJSON(true, constants.SUCCESS_CODE, constants.EMPTY_VALUE, userID)
	return ctx.JSON(http.StatusCreated, result)

}

func (h UserHandler) FindUserByID(ctx echo.Context) error {
	var result models.Response
	req := new(models.RequestID)
	if err := helpers.ValidateStruct(ctx, req); err != nil {
		log.Printf("Error Failed to validate request: %v", err)
		result = helpers.ResponseJSON(false, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	user, err := h.handler.UserService.FindUserByID(*req)
	if err != nil {
		log.Printf("Error FindUserByID: %v", err)
		result = helpers.ResponseJSON(false, constants.SYSTEM_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusInternalServerError, result)
	}
	result = helpers.ResponseJSON(true, constants.SUCCESS_CODE, constants.EMPTY_VALUE, user)
	return ctx.JSON(http.StatusOK, result)
}

func (h UserHandler) Login(ctx echo.Context) error {
	var result models.Response

	req := new(models.UserLoginRequest)
	if err := helpers.ValidateStruct(ctx, req); err != nil {
		log.Printf("Error Failed to validate request: %v", err)
		result = helpers.ResponseJSON(false, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	token, err := h.handler.UserService.Login(*req)
	if err != nil {
		log.Printf("Error Login: %v", err)
		result = helpers.ResponseJSON(false, constants.SYSTEM_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusInternalServerError, result)
	}
	result = helpers.ResponseJSON(true, constants.SUCCESS_CODE, constants.EMPTY_VALUE, token)
	return ctx.JSON(http.StatusOK, result)
}

func (h UserHandler) RefreshToken(ctx echo.Context) error {
	var result models.Response
	req := new(models.RefreshTokenRequest)
	if err := helpers.ValidateStruct(ctx, req); err != nil {
		log.Printf("Error Failed to validate request: %v", err)
		result = helpers.ResponseJSON(false, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	token, err := h.handler.UserService.RefreshToken(req.RefreshToken)
	if err != nil {
		log.Printf("Error RefreshToken: %v", err)
		result = helpers.ResponseJSON(false, constants.SYSTEM_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusInternalServerError, result)
	}
	result = helpers.ResponseJSON(true, constants.SUCCESS_CODE, constants.EMPTY_VALUE, token)
	return ctx.JSON(http.StatusOK, result)
}

func (h UserHandler) DeleteUser(ctx echo.Context) error {
	var result models.Response

	req := new(models.RequestID)
	if err := helpers.ValidateStruct(ctx, req); err != nil {
		log.Printf("Error Failed to validate request: %v", err)
		result = helpers.ResponseJSON(false, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	err := h.handler.UserService.DeleteUser(*req)
	if err != nil {
		log.Printf("Error DeleteUser: %v", err)
		result = helpers.ResponseJSON(false, constants.SYSTEM_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusInternalServerError, result)
	}

	result = helpers.ResponseJSON(true, constants.SUCCESS_CODE, constants.EMPTY_VALUE, "Successfully deleted")
	return ctx.JSON(http.StatusOK, result)
}

func (h UserHandler) FindListUsers(ctx echo.Context) error {
	var result models.Response
	req := new(models.FindListUserRequest)
	if err := helpers.ValidateStruct(ctx, req); err != nil {
		log.Printf("Error Failed to validate request: %v", err)
		result = helpers.ResponseJSON(false, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	list, err := h.handler.UserService.FindListUsers(*req)
	if err != nil {
		log.Printf("Error FindListUsers: %v", err)
		result = helpers.ResponseJSON(false, constants.SYSTEM_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusInternalServerError, result)
	}
	result = helpers.ResponseJSON(true, constants.SUCCESS_CODE, constants.EMPTY_VALUE, list)
	return ctx.JSON(http.StatusOK, result)
}

func (h UserHandler) CreateUser(ctx echo.Context) error {
	var result models.Response
	req := new(models.UserCreateRequest)
	if err := helpers.ValidateStruct(ctx, req); err != nil {
		log.Printf("Error Failed to validate request: %v", err)
		result = helpers.ResponseJSON(false, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	userID, err := h.handler.UserService.CreateUser(*req)
	if err != nil {
		log.Printf("Error CreateUser: %v", err)
		result = helpers.ResponseJSON(false, constants.SYSTEM_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusInternalServerError, result)
	}
	result = helpers.ResponseJSON(true, constants.SUCCESS_CODE, constants.EMPTY_VALUE, userID)
	return ctx.JSON(http.StatusCreated, result)
}

func (h UserHandler) CreateUserRolePermission(ctx echo.Context) error {
	var result models.Response
	req := new(models.UserRolePermissionCreateRequest)
	if err := helpers.ValidateStruct(ctx, req); err != nil {
		log.Printf("Error Failed to validate request: %v", err)
		result = helpers.ResponseJSON(false, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	permissionID, err := h.handler.UserPermissionService.CreateUserRolePermission(*req)
	if err != nil {
		log.Printf("Error CreateUserRolePermission: %v", err)
		result = helpers.ResponseJSON(false, constants.SYSTEM_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusInternalServerError, result)
	}
	result = helpers.ResponseJSON(true, constants.SUCCESS_CODE, constants.EMPTY_VALUE, permissionID)
	return ctx.JSON(http.StatusCreated, result)
}
