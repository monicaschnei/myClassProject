package request

import (
	"github.com/go-playground/validator/v10"
	"myclass/api"
)

type CreateProfessionalUserRequest struct {
	Name           string `json:"name" validate:"required"`
	Username       string `json:"username" validate:"required,alphanum"`
	Password       string `json:"password" validate:"required,min=8"`
	Gender         string `json:"gender" validate:"required"`
	Email          string `json:"email" validate:"required,email"`
	DateOfBirth    string `json:"date_of_birth"`
	Cpf            string `json:"cpf"`
	ClassHourPrice string `json:"class_hour_price"`
	ImageID        int64  `json:"image_id"`
}

type GetProfissionalUserRequest struct {
	UserName string `uri:"username" binding:"required"`
}

type LoginProfessionalUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

func (pur CreateProfessionalUserRequest) CreateProfessionalUserRequestValidator() (*validator.Validate, error) {
	validate := validator.New()
	if err := validate.RegisterValidation("passwd", api.ValidPassword); err != nil {
		return nil, err
	}
	if err := validate.RegisterValidation("gender", api.ValidGender); err != nil {
		return nil, err
	}
	return validate, nil
}
