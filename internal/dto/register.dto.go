package dto

type RegisterDTO struct {
	FirstName string `json:"first_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	Role      string `json:"role,omitempty" default:"user"`
}