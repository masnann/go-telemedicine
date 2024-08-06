package models

type RoleModels struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"isActive"`
}

type UserRoleModels struct {
	UserID int64 `json:"userID"`
	RoleID int64 `json:"roleID"`
}

type PermissionModels struct {
	ID        int64  `json:"id:"`
	Group     string `json:"group"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type RolePermissionModels struct {
	ID           int64  `json:"id:"`
	RoleID       int64  `json:"roleID"`
	PermissionID int64  `json:"permissionID"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

type UserPermissionModels struct {
	ID           int64  `json:"id:"`
	UserID       int64  `json:"userID"`
	PermissionID int64  `json:"permissionID"`
	Status       bool   `json:"status"`
	GrantedBy    int64  `json:"grantedBy"`
	GrantedAt    string `json:"grantedAt"`
	UpdatedAt    string `json:"updatedAt"`
}

type UserRolePermissionModels struct {
	ID     int64  `json:"id"`
	Group  string `json:"group"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

type AssignRoleToUserRequest struct {
	UserID int64 `json:"userID"`
	RoleID int64 `json:"roleID"`
}

type UserRolePermissionCreateRequest struct {
	RoleID       int64 `json:"roleID"`
	PermissionID int64 `json:"permissionID"`
}

type PermissionCreateRequest struct {
	Group string `json:"group"`
	Name  string `json:"name"`
}

type RolePermissionCreateRequest struct {
	RoleID       int64 `json:"roleID"`
	PermissionID int64 `json:"permissionID"`
}

type UserPermissionCreateRequest struct {
	AdminID      int64  `json:"adminID"`
	UserID       int64  `json:"userID"`
	PermissionID int64  `json:"permissionID"`
	Status       bool   `json:"status"`
	GrantedBy    int64  `json:"grantedBy"`
	GrantedAt    string `json:"grantedAt"`
	UpdatedAt    string `json:"updatedAt"`
}
