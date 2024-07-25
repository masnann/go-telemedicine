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

	req := new(models.ConsultationCreateRequest)
	req.PatientID = ctx.Get("user_id").(int64)
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
