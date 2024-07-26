package schedulehandler

import (
	"go-telemedicine/constants"
	"go-telemedicine/handler"
	"go-telemedicine/helpers"
	"go-telemedicine/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ScheduleHandler struct {
	handler handler.Handler
}

func NewScheduleHandler(handler handler.Handler) ScheduleHandler {
	return ScheduleHandler{
		handler: handler,
	}
}

func (h ScheduleHandler) CreateSchedule(ctx echo.Context) error {
	var result models.Response

	userID := ctx.Get("user_id").(int64)
	req := new(models.ScheduleCreateRequest)
	if err := helpers.ValidateStruct(ctx, req); err != nil {
		log.Printf("Error Failed to validate request: %v", err)
		result = helpers.ResponseJSON(false, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	req.DoctorID = userID
	scheduleID, err := h.handler.ScheduleService.CreateSchedule(*req)
	if err != nil {
		log.Printf("Error CreateSchedule: %v", err)
		result = helpers.ResponseJSON(false, constants.SYSTEM_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusInternalServerError, result)
	}
	result = helpers.ResponseJSON(true, constants.SUCCESS_CODE, constants.EMPTY_VALUE, scheduleID)
	return ctx.JSON(http.StatusOK, result)
}

func (h ScheduleHandler) FindListAvailableSchedule(ctx echo.Context) error {
	var result models.Response
	req := new(models.ScheduleFindListAvailableRequest)
	if err := helpers.ValidateStruct(ctx, req); err != nil {
		log.Printf("Error Failed to validate request: %v", err)
		result = helpers.ResponseJSON(false, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	schedules, err := h.handler.ScheduleService.FindListAvailableSchedule(*req)
	if err != nil {
		log.Printf("Error FindListAvailableSchedule: %v", err)
		result = helpers.ResponseJSON(false, constants.SYSTEM_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusInternalServerError, result)
	}
	result = helpers.ResponseJSON(true, constants.SUCCESS_CODE, constants.EMPTY_VALUE, schedules)
	return ctx.JSON(http.StatusOK, result)
}
