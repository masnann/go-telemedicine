package consultationservice

import (
	"errors"
	"go-telemedicine/helpers"
	"go-telemedicine/models"
	"go-telemedicine/service"
	"log"
)

type ConsultationService struct {
	service service.Service
}

func NewConsultationService(service service.Service) ConsultationService {
	return ConsultationService{
		service: service,
	}
}

func (s ConsultationService) CreateConsultation(req models.ConsultationCreateRequest) (int64, error) {
	patient, err := s.service.UserRepo.FindUserByID(req.PatientID)
	if err != nil {
		log.Println("Error finding patient by ID: ", err)
		return 0, errors.New("patient not found")
	}

	schedule, err := s.service.ScheduleRepo.FindScheduleByID(req.ScheduleID)
	if err != nil {
		log.Println("Error finding schedule by ID: ", err)
		return 0, errors.New("schedule not found")
	}

	if !schedule.IsAvailable {
		log.Println("Error: schedule is not available")
		return 0, errors.New("schedule is not available")
	}

	doctor, err := s.service.UserRepo.FindUserByID(schedule.DoctorID)
	if err != nil {
		log.Println("Error finding doctor by ID: ", err)
		return 0, errors.New("doctor not found")
	}

	newData := models.ConsultationModels{
		PatientID:        patient.ID,
		PatientName:      patient.Username,
		PatientEmail:     patient.Email,
		ScheduleID:       schedule.ID,
		DoctorID:         schedule.DoctorID,
		DoctorName:       doctor.Username,
		StartTime:        schedule.StartTime,
		EndTime:          schedule.EndTime,
		ConsultationType: req.ConsultationType,
		Status:           "Scheduled",
		Notes:            req.Notes,
		CreatedAt:        helpers.TimeStampNow(),
	}
	result, err := s.service.ConsultationRepo.CreateConsultation(newData)
	if err != nil {
		return 0, err
	}
	return result, nil
}
