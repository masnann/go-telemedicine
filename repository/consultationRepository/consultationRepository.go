package consultationrepository

import (
	"errors"
	"fmt"
	"go-telemedicine/constants"
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

func (r ConsultationRepository) FindListConsultationsUser(req models.ConsultationFindListByPatientIDRequest, userType string) ([]models.ConsultationModels, error) {
	var list []models.ConsultationModels
	var params []interface{}

	query := fmt.Sprintf(`
		SELECT
			id, 
			patient_id, 
			patient_name,
			patient_email, 
			schedule_id, 
			doctor_id, 
			doctor_name, 
			start_time, 
			end_time, 
			consultation_type, 
			status, 
			notes, 
			created_at, 
			updated_at
		FROM 
			consultations
		WHERE 
			%s_id = ? AND deleted_at =''
	`, userType)

	params = append(params, req.PatientID)

	if req.DateStart != constants.EMPTY_STRING || req.DateEnd != constants.EMPTY_STRING {
		if req.DateStart != constants.EMPTY_STRING && req.DateEnd == constants.EMPTY_STRING {
			query += ` AND start_time >= ?`
			params = append(params, req.DateStart)
		} else if req.DateStart == constants.EMPTY_STRING && req.DateEnd != constants.EMPTY_STRING {
			query += ` AND end_time <= ?`
			params = append(params, req.DateEnd)
		} else if req.DateStart != constants.EMPTY_STRING && req.DateEnd != constants.EMPTY_STRING {
			query += ` AND start_time BETWEEN ? AND ?`
			params = append(params, req.DateStart, req.DateEnd)
		}
	}
	if req.Status != constants.EMPTY_STRING {
		query += ` AND status = ?`
		params = append(params, req.Status)
	}

	query += ` ORDER BY created_at DESC`
	if req.Pagination.Page != constants.EMPTY_INT && req.Pagination.PageSize != constants.EMPTY_INT {
		offset := (req.Pagination.Page - 1) * req.Pagination.PageSize
		query += ` LIMIT ? OFFSET ?`
		params = append(params, req.Pagination.PageSize, offset)
	}

	query = helpers.ReplaceSQL(query, "?")

	rows, err := r.repo.DB.Query(query, params...)
	if err != nil {
		log.Println("Error querying find list consultations by patient id: ", err)
		return list, errors.New("error query")
	}
	defer rows.Close()

	for rows.Next() {
		var row models.ConsultationModels
		err := rows.Scan(
			&row.ID,
			&row.PatientID,
			&row.PatientName,
			&row.PatientEmail,
			&row.ScheduleID,
			&row.DoctorID,
			&row.DoctorName,
			&row.StartTime,
			&row.EndTime,
			&row.ConsultationType,
			&row.Status,
			&row.Notes,
			&row.CreatedAt,
			&row.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning row: ", err)
			return list, errors.New("error scanning row")
		}
		list = append(list, row)
	}
	return list, nil
}
