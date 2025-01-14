package dto

type UserFilter struct {
	Name  string   `form:"name"`
	Email string   `form:"email"`
	Role  []string `form:"role[]"`
}
