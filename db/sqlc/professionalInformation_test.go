package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateProfessionalInformation(t *testing.T) {
	argUser := CreateProfessionalUserParams{
		ProfessionalInformationID: 123,
		Name:                      "Monica de Souza Schneider",
		Username:                  "profmonica",
		Password:                  "profmonica",
		Gender:                    "feminino",
		Email:                     "profmonica@gmail.com",
		DateOfBirth:               time.Date(1992, 01, 13, 0, 0, 0, 0, time.UTC),
		Cpf:                       11111111,
	}
	argInfo := CreateProfessionalInformationParams{
		ID:                123,
		ExperiencePeriod:  "2 years",
		OcupationArea:     "Math",
		University:        "UDESC",
		GraduationDiploma: "Diploma",
		Validate:          false,
		GraduationCountry: "Brasil",
		GraduationCity:    "Joinvile",
		GraduationState:   "SC",
		UpdatedAt:         time.Date(2022, 12, 1, 0, 0, 0, 0, time.UTC),
	}
	_, err := testQueries.CreateProfessionalUser(context.Background(), argUser)
	_, err = testQueries.CreateProfessionalInformation(context.Background(), argInfo)

	assert.NoError(t, err)
	//assert.NotEmpty(t, professionalUser)

}
