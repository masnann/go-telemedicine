package schedulerepository

import (
	"database/sql"
	"errors"
	"go-telemedicine/constants"
	"go-telemedicine/helpers"
	"go-telemedicine/models"
	"go-telemedicine/repository"
	"log"
)

type ScheduleRepository struct {
	repo repository.Repository
}

func NewScheduleRepository(repo repository.Repository) ScheduleRepository {
	return ScheduleRepository{
		repo: repo,
	}
}

func (r ScheduleRepository) CreateSchedule(req models.ScheduleModels) (int64, error) {
	var ID int64
	query := `
        INSERT INTO schedules 
			(doctor_id, date, start_time, end_time, is_available, created_at, updated_at) 
        VALUES 
			(?,?,?,?,?,?,?)
        RETURNING id`

	query = helpers.ReplaceSQL(query, "?")
	err := r.repo.DB.QueryRow(query, req.DoctorID, req.Date, req.StartTime, req.EndTime, req.IsAvailable, req.CreatedAt, req.UpdatedAt).Scan(&ID)
	if err != nil {
		log.Println("Error querying create schedule: ", err)
		return ID, errors.New("error query")
	}

	return ID, err
}

func (r ScheduleRepository) FindListAvailableSchedule(req models.ScheduleFindListAvailableRequest) ([]models.ScheduleModels, error) {
	var schedules []models.ScheduleModels
	var params []interface{}

	query := `
        SELECT 
			id, 
			doctor_id, 
			date, 
			start_time, 
			end_time, 
			is_available
        FROM schedules
        WHERE deleted_at IS NULL`

	if req.Date != constants.EMPTY_STRING {
		query += ` AND DATE = ?`
		params = append(params, req.Date)
	}
	if req.StartTime != constants.EMPTY_STRING && req.EndTime != constants.EMPTY_STRING {
		query += ` AND (start_time BETWEEN ? AND ? OR end_time BETWEEN ? AND ?)`
		params = append(params, req.StartTime, req.EndTime, req.StartTime, req.EndTime)
	}
	query = helpers.ReplaceSQL(query, "?")

	rows, err := r.repo.DB.Query(query, params...)
	if err != nil {
		log.Println("Error querying schedules: ", err)
		return schedules, errors.New("error query")
	}
	defer rows.Close()

	for rows.Next() {
		var row models.ScheduleModels
		err := rows.Scan(&row.ID, &row.DoctorID, &row.Date, &row.StartTime, &row.EndTime, &row.IsAvailable)
		if err != nil {
			log.Println("Error scanning row: ", err)
			return schedules, errors.New("error scanning row")
		}
		schedules = append(schedules, row)
	}
	return schedules, nil

}

func (s ScheduleRepository) FindScheduleByID(id int64) (models.ScheduleModels, error) {
	var schedules models.ScheduleModels

	query := `
		SELECT 
			id, 
			doctor_id, 
			date, 
			start_time, 
			end_time, 
			is_available, 
			created_at, updated_at
		FROM schedules WHERE id = ? AND deleted_at is NULL `

	query = helpers.ReplaceSQL(query, "?")
	row := s.repo.DB.QueryRow(query, id)
	err := row.Scan(
		&schedules.ID, &schedules.DoctorID, &schedules.Date,
		&schedules.StartTime, &schedules.EndTime, &schedules.IsAvailable,
		&schedules.CreatedAt, &schedules.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return schedules, errors.New("schedule not found")
		}
		log.Println("Error scanning row: ", err)
		return schedules, errors.New("error scanning row")
	}
	return schedules, nil
}