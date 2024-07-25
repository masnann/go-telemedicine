package consultationrepository

import (
	"errors"
	"go-telemedicine/helpers"
	"go-telemedicine/models"
	"go-telemedicine/repository"
	"log"
)

type ConsultationRepository struct {
	repo repository.Repository
}

func NewConsultationRepository(repo repository.Repository) ConsultationRepository {
	return ConsultationRepository{
		repo: repo,
	}
}

func (r ConsultationRepository) CreateConsultation(req models.ConsultationModels) (int64, error) {
	var ID int64

	query := ` 
		INSERT INTO consultations
		    (patient_id, patient_name, patient_email, schedule_id, doctor_id, doctor_name, 
				start_time, end_time, consultation_type, status, notes, created_at, updated_at, deleted_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?)
		RETURNING id	
	`
	query = helpers.ReplaceSQL(query, "?")
	err := r.repo.DB.QueryRow(query,
        req.PatientID, req.PatientName,
        req.PatientEmail,
        req.ScheduleID,
		req.DoctorID,
        req.DoctorName,
        req.StartTime,
        req.EndTime,
        req.ConsultationType,
		req.Status,
        req.Notes,
        req.CreatedAt,
        req.UpdatedAt,
        req.DeletedAt).Scan(&ID)
	if err != nil {
		log.Println("Error querying create consultation: ", err)
		return ID, errors.New("error query")
	}
	return ID, nil

}
