package models

type UserModels struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UserRegisterRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	UserID       int64                      `json:"userID"`
	RoleName     string                     `json:"roleName"`
	AccessToken  string                     `json:"accessToken"`
	RefreshToken string                     `json:"refreshToken"`
	Permission   []UserRolePermissionModels `json:"permission"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type FindUserRoleResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	RoleID   int    `json:"roleID"`
	RoleName string `json:"roleName"`
}

type FindListUserRequest struct {
	Status     string     `json:"status"`
	Pagination Pagination `json:"pagination"`
}

type FindListUserResponse struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
}

type UserCreateRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Status    string `json:"status"`
	RoleID    int64  `json:"roleID"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type CurrentUserModels struct {
	ID    int64  `json:"id"`
	Role  string `json:"role"`
	Email string `json:"email"`
}
