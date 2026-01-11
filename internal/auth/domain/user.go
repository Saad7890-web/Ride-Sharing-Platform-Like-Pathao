package domain

import "time"

type User struct {
	ID string
	Email string
	PasswordHash string
	IsActive bool
	Roles []Role
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) HasPermission(permission string) bool {
	for _, role := range u.Roles {
		for _, perm := range role.Permissions {
			if perm.Name == permission {
				return true
			}
		}
	}
	return false
}