package userrepository

import (
	"database/sql"
	"errors"
	"go-telemedicine/constants"
	"go-telemedicine/helpers"
	"go-telemedicine/models"
	"go-telemedicine/repository"
	"log"
)

type UserRepository struct {
	repo repository.Repository
}

func NewUserRepository(repo repository.Repository) UserRepository {
	return UserRepository{
		repo: repo,
	}
}

func (r UserRepository) Register(req models.UserModels) (int64, error) {
	var ID int64
	query := `
		INSERT INTO users (username, email, password, status, created_at, updated_at) 
		VALUES (?, ?, ?, ?, ?, ?)
		RETURNING id`

	query = helpers.ReplaceSQL(query, "?")
	err := r.repo.DB.QueryRow(query, req.Username, req.Email, req.Password, req.Status, req.CreatedAt, req.UpdatedAt).Scan(&ID)
	if err != nil {
		log.Println("Error querying register: ", err)
		return ID, errors.New("error query")
	}

	return ID, nil
}

func (r UserRepository) FindUserByID(id int64) (models.UserModels, error) {
	var user models.UserModels

	query := `SELECT * FROM users WHERE id = ? and status = 'active'`

	query = helpers.ReplaceSQL(query, "?")
	row := r.repo.DB.QueryRow(query, id)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Status, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("user not found")
		}
		log.Println("Error scanning row: ", err)
		return user, errors.New("error scanning row")
	}

	return user, nil
}

func (r UserRepository) Login(email string) (models.UserModels, error) {
	var user models.UserModels
	query := `
		SELECT 
			id, 
			username, 
			email, 
			password
		FROM users 
		WHERE email =?
	`

	query = helpers.ReplaceSQL(query, "?")

	rows, err := r.repo.DB.Query(query, email)
	if err != nil {
		log.Println("Error querying login: ", err)
		return user, errors.New("error query")
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
		if err != nil {
			log.Println("Error scanning row: ", err)
			return user, errors.New("error scanning row")
		}
	}

	return user, nil
}

func (r UserRepository) DeleteUser(userID int64) error {
	query := `
		UPDATE 
			users 
		SET 
			status = ? 
		WHERE 
			id = ?
	`
	query = helpers.ReplaceSQL(query, "?")
	_, err := r.repo.DB.Exec(query, constants.USER_STATUS_DELETE, userID)
	if err != nil {
		log.Println("Error querying delete user: ", err)
		return errors.New("error query")
	}
	return nil
}

func (r UserRepository) FindListUser(req models.FindListUserRequest) ([]models.FindListUserResponse, error) {
	var user []models.FindListUserResponse
	var params []interface{}

	query := `
		SELECT
		    id, 
            username, 
            email, 
            status, 
            created_at
		FROM 
			users
		WHERE true
		
	`
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
		log.Println("Error querying find list user: ", err)
		return user, errors.New("error query")
	}
	defer rows.Close()
	for rows.Next() {
		var row models.FindListUserResponse
		err := rows.Scan(&row.ID, &row.Username, &row.Email, &row.Status, &row.CreatedAt)
		if err != nil {
			log.Println("Error scanning row: ", err)
			return user, errors.New("error scanning row")
		}
		user = append(user, row)
	}
	return user, nil
}
