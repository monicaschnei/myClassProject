// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: professionalInformation.sql

package db

import (
	"context"
	"time"
)

const createProfessionalInformation = `-- name: CreateProfessionalInformation :one
INSERT INTO "professionalInformation" (
  professional_user_id,
  experience_period,
  ocupation_area,
  university,
  graduation_diploma,
  validate,
  graduation_country,
  graduation_city,
  graduation_state,
  updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9,$10
)
RETURNING id, experience_period, ocupation_area, university, graduation_diploma, validate, graduation_country, graduation_city, graduation_state, created_at, updated_at, professional_user_id
`

type CreateProfessionalInformationParams struct {
	ProfessionalUserID int64     `json:"professional_user_id"`
	ExperiencePeriod   string    `json:"experience_period"`
	OcupationArea      string    `json:"ocupation_area"`
	University         string    `json:"university"`
	GraduationDiploma  string    `json:"graduation_diploma"`
	Validate           bool      `json:"validate"`
	GraduationCountry  string    `json:"graduation_country"`
	GraduationCity     string    `json:"graduation_city"`
	GraduationState    string    `json:"graduation_state"`
	UpdatedAt          time.Time `json:"updated_at"`
}

func (q *Queries) CreateProfessionalInformation(ctx context.Context, arg CreateProfessionalInformationParams) (ProfessionalInformation, error) {
	row := q.queryRow(ctx, q.createProfessionalInformationStmt, createProfessionalInformation,
		arg.ProfessionalUserID,
		arg.ExperiencePeriod,
		arg.OcupationArea,
		arg.University,
		arg.GraduationDiploma,
		arg.Validate,
		arg.GraduationCountry,
		arg.GraduationCity,
		arg.GraduationState,
		arg.UpdatedAt,
	)
	var i ProfessionalInformation
	err := row.Scan(
		&i.ID,
		&i.ExperiencePeriod,
		&i.OcupationArea,
		&i.University,
		&i.GraduationDiploma,
		&i.Validate,
		&i.GraduationCountry,
		&i.GraduationCity,
		&i.GraduationState,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProfessionalUserID,
	)
	return i, err
}

const deleteProfessionalInformation = `-- name: DeleteProfessionalInformation :one
DELETE FROM "professionalInformation"
WHERE id = $1
RETURNING id, experience_period, ocupation_area, university, graduation_diploma, validate, graduation_country, graduation_city, graduation_state, created_at, updated_at, professional_user_id
`

func (q *Queries) DeleteProfessionalInformation(ctx context.Context, id int64) (ProfessionalInformation, error) {
	row := q.queryRow(ctx, q.deleteProfessionalInformationStmt, deleteProfessionalInformation, id)
	var i ProfessionalInformation
	err := row.Scan(
		&i.ID,
		&i.ExperiencePeriod,
		&i.OcupationArea,
		&i.University,
		&i.GraduationDiploma,
		&i.Validate,
		&i.GraduationCountry,
		&i.GraduationCity,
		&i.GraduationState,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProfessionalUserID,
	)
	return i, err
}

const getProfessionalInformation = `-- name: GetProfessionalInformation :one
SELECT id, experience_period, ocupation_area, university, graduation_diploma, validate, graduation_country, graduation_city, graduation_state, created_at, updated_at, professional_user_id FROM "professionalInformation" 
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetProfessionalInformation(ctx context.Context, id int64) (ProfessionalInformation, error) {
	row := q.queryRow(ctx, q.getProfessionalInformationStmt, getProfessionalInformation, id)
	var i ProfessionalInformation
	err := row.Scan(
		&i.ID,
		&i.ExperiencePeriod,
		&i.OcupationArea,
		&i.University,
		&i.GraduationDiploma,
		&i.Validate,
		&i.GraduationCountry,
		&i.GraduationCity,
		&i.GraduationState,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProfessionalUserID,
	)
	return i, err
}

const listProfessionalInformation = `-- name: ListProfessionalInformation :many
SELECT id, experience_period, ocupation_area, university, graduation_diploma, validate, graduation_country, graduation_city, graduation_state, created_at, updated_at, professional_user_id FROM "professionalInformation" 
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListProfessionalInformationParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListProfessionalInformation(ctx context.Context, arg ListProfessionalInformationParams) ([]ProfessionalInformation, error) {
	rows, err := q.query(ctx, q.listProfessionalInformationStmt, listProfessionalInformation, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ProfessionalInformation{}
	for rows.Next() {
		var i ProfessionalInformation
		if err := rows.Scan(
			&i.ID,
			&i.ExperiencePeriod,
			&i.OcupationArea,
			&i.University,
			&i.GraduationDiploma,
			&i.Validate,
			&i.GraduationCountry,
			&i.GraduationCity,
			&i.GraduationState,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ProfessionalUserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listProfessionalInformationByUser = `-- name: ListProfessionalInformationByUser :many
SELECT id, experience_period, ocupation_area, university, graduation_diploma, validate, graduation_country, graduation_city, graduation_state, created_at, updated_at, professional_user_id FROM "professionalInformation"
WHERE professional_user_id = $1
LIMIT $1
OFFSET $2
`

type ListProfessionalInformationByUserParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListProfessionalInformationByUser(ctx context.Context, arg ListProfessionalInformationByUserParams) ([]ProfessionalInformation, error) {
	rows, err := q.query(ctx, q.listProfessionalInformationByUserStmt, listProfessionalInformationByUser, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ProfessionalInformation{}
	for rows.Next() {
		var i ProfessionalInformation
		if err := rows.Scan(
			&i.ID,
			&i.ExperiencePeriod,
			&i.OcupationArea,
			&i.University,
			&i.GraduationDiploma,
			&i.Validate,
			&i.GraduationCountry,
			&i.GraduationCity,
			&i.GraduationState,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ProfessionalUserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProfessionalInformation = `-- name: UpdateProfessionalInformation :one
UPDATE "professionalInformation"
    set graduation_state = $2,
    experience_period = $3,
    ocupation_area = $4,
    university = $5,
    graduation_diploma = $6,
    validate = $7,
    graduation_country = $8,
    graduation_city = $9
WHERE id = $1
RETURNING id, experience_period, ocupation_area, university, graduation_diploma, validate, graduation_country, graduation_city, graduation_state, created_at, updated_at, professional_user_id
`

type UpdateProfessionalInformationParams struct {
	ID                int64  `json:"id"`
	GraduationState   string `json:"graduation_state"`
	ExperiencePeriod  string `json:"experience_period"`
	OcupationArea     string `json:"ocupation_area"`
	University        string `json:"university"`
	GraduationDiploma string `json:"graduation_diploma"`
	Validate          bool   `json:"validate"`
	GraduationCountry string `json:"graduation_country"`
	GraduationCity    string `json:"graduation_city"`
}

func (q *Queries) UpdateProfessionalInformation(ctx context.Context, arg UpdateProfessionalInformationParams) (ProfessionalInformation, error) {
	row := q.queryRow(ctx, q.updateProfessionalInformationStmt, updateProfessionalInformation,
		arg.ID,
		arg.GraduationState,
		arg.ExperiencePeriod,
		arg.OcupationArea,
		arg.University,
		arg.GraduationDiploma,
		arg.Validate,
		arg.GraduationCountry,
		arg.GraduationCity,
	)
	var i ProfessionalInformation
	err := row.Scan(
		&i.ID,
		&i.ExperiencePeriod,
		&i.OcupationArea,
		&i.University,
		&i.GraduationDiploma,
		&i.Validate,
		&i.GraduationCountry,
		&i.GraduationCity,
		&i.GraduationState,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProfessionalUserID,
	)
	return i, err
}
