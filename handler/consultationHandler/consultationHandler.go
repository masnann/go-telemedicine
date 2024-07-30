package consultationhandler

import (
	"go-telemedicine/constants"
	"go-telemedicine/handler"
	"go-telemedicine/helpers"
	"go-telemedicine/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ConsultationHandler struct {
	handler handler.Handler
}

func NewConsultationHandler(handler handler.Handler) ConsultationHandler {
	return ConsultationHandler{
		handler: handler,
	}
}

func (h ConsultationHandler) CreateConsultation(ctx echo.Context) error {
	var result models.Response

	currentUser, ok := ctx.Get("user").(models.CurrentUserModels)
	if !ok {
		result = helpers.ResponseJSON(false, constants.UNAUTHORIZED_CODE, "Failed to get user from context", nil)
		return ctx.JSON(http.StatusInternalServerError, result)
	}

	if currentUser.Role != "Admin" && currentUser.Role != "Patient" {
		result = helpers.ResponseJSON(false, constants.FORBIDDEN_CODE, "Access denied. You don't have permission", nil)
		return ctx.JSON(http.StatusForbidden, result)
	}

	req := new(models.ConsultationCreateRequest)
	req.PatientID = currentUser.ID
	if err := helpers.ValidateStruct(ctx, req); err != nil {
		log.Printf("Error Failed to validate request: %v", err)
		result = helpers.ResponseJSON(false, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	log.Println(req.PatientID)
	consultationID, err := h.handler.ConsultationService.CreateConsultation(*req)
	if err != nil {
		log.Printf("Error CreateConsultation: %v", err)
		result = helpers.ResponseJSON(false, constants.SYSTEM_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusInternalServerError, result)
	}
	result = helpers.ResponseJSON(true, constants.SUCCESS_CODE, constants.EMPTY_VALUE, consultationID)
	return ctx.JSON(http.StatusOK, result)
}

func (h ConsultationHandler) FindListConsultationsByPatientID(ctx echo.Context) error {
	var result models.Response

	req := new(models.ConsultationFindListByPatientIDRequest)
	req.PatientID = ctx.Get("user_id").(int64)
	if err := helpers.ValidateStruct(ctx, req); err != nil {
		log.Printf("Error Failed to validate request: %v", err)
		result = helpers.ResponseJSON(false, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	consultations, err := h.handler.ConsultationService.FindListConsultationsByPatientID(*req)
	if err != nil {
		log.Printf("Error FindListConsultationsByPatientID: %v", err)
		result = helpers.ResponseJSON(false, constants.SYSTEM_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusInternalServerError, result)
	}
	if len(consultations) == 0 {
		result = helpers.ResponseJSON(false, constants.DATA_NOT_FOUND_CODE, "consultations not found", nil)
		return ctx.JSON(http.StatusNotFound, result)
	}

	result = helpers.ResponseJSON(true, constants.SUCCESS_CODE, constants.EMPTY_VALUE, consultations)
	return ctx.JSON(http.StatusOK, result)
}

func (h ConsultationHandler) FindListConsultationsByDoctorID(ctx echo.Context) error {
	var result models.Response

	req := new(models.ConsultationFindListByDoctorIDRequest)
	req.DoctorID = ctx.Get("user_id").(int64)
	if err := helpers.ValidateStruct(ctx, req); err != nil {
		log.Printf("Error Failed to validate request: %v", err)
		result = helpers.ResponseJSON(false, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	consultations, err := h.handler.ConsultationService.FindListConsultationsByDoctorID(*req)
	if err != nil {
		log.Printf("Error FindListConsultationsByDoctorID: %v", err)
		result = helpers.ResponseJSON(false, constants.SYSTEM_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusInternalServerError, result)
	}
	if len(consultations) == 0 {
		result = helpers.ResponseJSON(false, constants.DATA_NOT_FOUND_CODE, "consultations not found", nil)
		return ctx.JSON(http.StatusNotFound, result)
	}
	result = helpers.ResponseJSON(true, constants.SUCCESS_CODE, constants.EMPTY_VALUE, consultations)
	return ctx.JSON(http.StatusOK, result)
}
