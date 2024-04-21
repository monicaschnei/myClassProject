package user

import (
	db "myclass/db/sqlc"
)

type UserSession struct {
	User        *db.ProfessionalUser
	TokenAccess string
}
