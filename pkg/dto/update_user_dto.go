package dto

type UpdateUserRequest struct {
	Name string `json:"name" binding:"omitempty,min=3,max=30"`
}
