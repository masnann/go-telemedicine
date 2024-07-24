package userrepository

import (
	"errors"
	"go-telemedicine/helpers"
	"go-telemedicine/models"
	"go-telemedicine/repository"
	"log"
)

type UserPermissionRepository struct {
	repo repository.Repository
}

func NewUserPermissionRepository(repo repository.Repository) UserPermissionRepository {
	return UserPermissionRepository{
		repo: repo,
	}
}

func (r UserPermissionRepository) FindListUserPermissions(userID int64) ([]models.UserPermissionModels, error) {
	var permissions []models.UserPermissionModels
	query := `
		SELECT 
			p.id AS permissions_id,
			p.groups AS permission_group,
			p.name AS permission_name,
			rp.status
		FROM 
			users u
		JOIN 
			user_role ur ON u.id = ur.user_id
		JOIN 
			roles r ON ur.role_id = r.id
		JOIN 
			role_permissions rp ON r.id = rp.role_id
		JOIN 
			permissions p ON rp.permission_id = p.id
		WHERE 
			u.id = ?

    `
	query = helpers.ReplaceSQL(query, "?")

	rows, err := r.repo.DB.Query(query, userID)
	if err != nil {
		log.Println("Error querying find user permission: ", err)
		return permissions, errors.New("error query")
	}
	defer rows.Close()
	for rows.Next() {
		var permission models.UserPermissionModels
		err := rows.Scan(&permission.ID, &permission.Group, &permission.Name, &permission.Status)
		if err != nil {
			log.Println("Error scanning row: ", err)
			return permissions, errors.New("error scanning row")
		}
		permissions = append(permissions, permission)
	}
	return permissions, nil
}

func (r UserPermissionRepository) AssignRoleToUserRequest(req models.AssignRoleToUserRequest) error {

	query := `
        INSERT INTO user_role (user_id, role_id) 
        VALUES (?,?)
	`

	query = helpers.ReplaceSQL(query, "?")
	_, err := r.repo.DB.Exec(query, req.UserID, req.RoleID)
	if err != nil {
		log.Println("Error querying create user role: ", err)
		return errors.New("error query")
	}

	return nil
}

func (r UserPermissionRepository) FindUserRole(userID int64) (models.FindUserRoleResponse, error) {
	var userRole models.FindUserRoleResponse
	query := ` 
		SELECT
			u.username,
			u.email,
			r.id as role_id, 
			r.name as role_name
		FROM 
			users u
		JOIN 
			user_role ur ON u.id = ur.user_id
		JOIN 
			roles r ON ur.role_id = r.id 
		WHERE
			u.id = ?
		`
	query = helpers.ReplaceSQL(query, "?")
	rows, err := r.repo.DB.Query(query, userID)
	if err != nil {
		log.Println("Error querying find user role: ", err)
		return userRole, errors.New("error query")
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&userRole.Username, &userRole.Email, &userRole.RoleID, &userRole.RoleName)
		if err != nil {
			log.Println("Error scanning row: ", err)
			return userRole, errors.New("error scanning row")
		}
	}
	return userRole, nil
}

func (r UserPermissionRepository) FindUserPermissions(userID int64, permissionGroup, permissionName string) (models.UserPermissionModels, error) {
	var permission models.UserPermissionModels
	query := `
        SELECT 
            p.id AS permissions_id,
            p.groups AS permission_group,
            p.name AS permission_name,
            rp.status
        FROM 
            users u
        JOIN 
            user_role ur ON u.id = ur.user_id
        JOIN 
            roles r ON ur.role_id = r.id
        JOIN 
            role_permissions rp ON r.id = rp.role_id
        JOIN 
            permissions p ON rp.permission_id = p.id
        WHERE 
            u.id =? AND p.groups =? AND p.name =?
    `

	query = helpers.ReplaceSQL(query, "?")

	rows, err := r.repo.DB.Query(query, userID, permissionGroup, permissionName)
	if err != nil {
		log.Println("Error querying find user permission: ", err)
		return permission, errors.New("error query")
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&permission.ID, &permission.Group, &permission.Name, &permission.Status)
		if err != nil {
			log.Println("Error scanning row: ", err)
			return permission, errors.New("error scanning row")
		}
	}
	return permission, nil
}
