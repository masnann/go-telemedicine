package scheduleservice

import (
	"errors"
	"go-telemedicine/helpers"
	"go-telemedicine/models"
	"go-telemedicine/service"
	"log"
)

type ScheduleService struct {
	service service.Service
}

func NewScheduleService(service service.Service) ScheduleService {
	return ScheduleService{
		service: service,
	}
}

func (s ScheduleService) CreateSchedule(req models.ScheduleCreateRequest) (int64, error) {
	newData := models.ScheduleModels{
		DoctorID:    req.DoctorID,
		Date:        req.Date,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		IsAvailable: true,
		CreatedAt:   helpers.TimeStampNow(),
		UpdatedAt:   "",
	}
	result, err := s.service.ScheduleRepo.CreateSchedule(newData)
	if err != nil {
		log.Println("Error creating schedule: ", err)
		return 0, err
	}
	return result, nil
}

func (s ScheduleService) FindListAvailableSchedule(req models.ScheduleFindListAvailableRequest) ([]models.ScheduleModels, error) {
	result, err := s.service.ScheduleRepo.FindListAvailableSchedule(req)
	if err != nil {
		log.Println("Error finding available schedules: ", err)
		return nil, err
	}
	return result, nil
}

func (s ScheduleService) FindByID(req models.RequestID) (models.ScheduleModels, error) {
	result, err := s.service.ScheduleRepo.FindScheduleByID(req.ID)
	if err != nil {
		log.Println("Error finding schedule by ID: ", err)
		return result, errors.New("schedule not found")
	}
	return result, nil
}
