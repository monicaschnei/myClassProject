package response

import (
	db "myclass/db/sqlc"
	"time"
)

type ProfessionalUserResponse struct {
	Name              string    `json:"name"`
	Username          string    `json:"username"`
	Gender            string    `json:"gender"`
	Email             string    `json:"email"`
	DateOfBirth       time.Time `json:"date_of_birth"`
	Cpf               string    `json:"cpf"`
	ClassHourPrice    string    `json:"class_hour_price"`
	ImageID           int64     `json:"image_id"`
	CreatedAt         time.Time `json:"createdAt"`
	PasswordChangedAt time.Time `json:"passwordChangedAt"`
}

type LoginProfessionalUserResponse struct {
	AccessToken string                   `json:"access_token"`
	User        ProfessionalUserResponse `json:"user"`
}

func NewProfessionalUserResponse(professionalUser db.ProfessionalUser) ProfessionalUserResponse {
	return ProfessionalUserResponse{
		Name:              professionalUser.Name,
		Username:          professionalUser.Username,
		Gender:            professionalUser.Gender,
		Email:             professionalUser.Email,
		DateOfBirth:       professionalUser.DateOfBirth,
		Cpf:               professionalUser.Cpf,
		ClassHourPrice:    professionalUser.ClassHourPrice,
		ImageID:           professionalUser.ImageID,
		CreatedAt:         professionalUser.CreatedAt,
		PasswordChangedAt: professionalUser.PasswordChangedAt,
	}
}
