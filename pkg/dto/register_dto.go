package dto

type RegisterRequest struct {
	Name            string   `json:"name" binding:"required,min=3,max=30"`
	Email           string   `json:"email" binding:"required,email"`
	Password        string   `json:"password" binding:"required,min=6,password_validator"`
	Roles           []string `json:"roles" binding:"omitempty"`
	ConfirmPassword string   `json:"confirm_password" binding:"required,eqfield=Password"`
}
