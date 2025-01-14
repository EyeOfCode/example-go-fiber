package utils

type Role string

const (
	UserRole  Role = "user"
	AdminRole Role = "admin"
)

func IsValidRole(r []Role, reqRole []Role) bool {
	hasRequiredRole := false
	for _, requiredRole := range reqRole {
		for _, userRole := range r {
			if Role(userRole) == requiredRole {
				hasRequiredRole = true
				return hasRequiredRole
			}
		}
	}
	return hasRequiredRole
}
