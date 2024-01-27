// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: professionalUser.sql

package db

import (
	"context"
	"time"
)

const createProfessionalUser = `-- name: CreateProfessionalUser :one
INSERT INTO "professionalUser" (
    name,
    username,
    password,
    gender,
    email, 
    date_of_birth,
    cpf,
    image_id,
    phone_id,
    professional_information_id,
    updated_at,
    "subjectMatter_id",
    "subjectMatter_class_id",
    class_hour_price,
    calendar_id 
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15 
)
RETURNING id, created_at, name, username, password, gender, email, date_of_birth, cpf, image_id, phone_id, professional_information_id, updated_at, "subjectMatter_id", "subjectMatter_class_id", class_hour_price, calendar_id
`

type CreateProfessionalUserParams struct {
	Name                      string    `json:"name"`
	Username                  string    `json:"username"`
	Password                  string    `json:"password"`
	Gender                    string    `json:"gender"`
	Email                     string    `json:"email"`
	DateOfBirth               time.Time `json:"date_of_birth"`
	Cpf                       int32     `json:"cpf"`
	ImageID                   int64     `json:"image_id"`
	PhoneID                   int64     `json:"phone_id"`
	ProfessionalInformationID int64     `json:"professional_information_id"`
	UpdatedAt                 time.Time `json:"updated_at"`
	SubjectMatterID           int32     `json:"subjectMatter_id"`
	SubjectMatterClassID      int32     `json:"subjectMatter_class_id"`
	ClassHourPrice            string    `json:"class_hour_price"`
	CalendarID                int32     `json:"calendar_id"`
}

func (q *Queries) CreateProfessionalUser(ctx context.Context, arg CreateProfessionalUserParams) (ProfessionalUser, error) {
	row := q.queryRow(ctx, q.createProfessionalUserStmt, createProfessionalUser,
		arg.Name,
		arg.Username,
		arg.Password,
		arg.Gender,
		arg.Email,
		arg.DateOfBirth,
		arg.Cpf,
		arg.ImageID,
		arg.PhoneID,
		arg.ProfessionalInformationID,
		arg.UpdatedAt,
		arg.SubjectMatterID,
		arg.SubjectMatterClassID,
		arg.ClassHourPrice,
		arg.CalendarID,
	)
	var i ProfessionalUser
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Name,
		&i.Username,
		&i.Password,
		&i.Gender,
		&i.Email,
		&i.DateOfBirth,
		&i.Cpf,
		&i.ImageID,
		&i.PhoneID,
		&i.ProfessionalInformationID,
		&i.UpdatedAt,
		&i.SubjectMatterID,
		&i.SubjectMatterClassID,
		&i.ClassHourPrice,
		&i.CalendarID,
	)
	return i, err
}

const deleteProfessionalUser = `-- name: DeleteProfessionalUser :one
DELETE FROM "professionalUser"
WHERE id = $1
RETURNING id, created_at, name, username, password, gender, email, date_of_birth, cpf, image_id, phone_id, professional_information_id, updated_at, "subjectMatter_id", "subjectMatter_class_id", class_hour_price, calendar_id
`

func (q *Queries) DeleteProfessionalUser(ctx context.Context, id int64) (ProfessionalUser, error) {
	row := q.queryRow(ctx, q.deleteProfessionalUserStmt, deleteProfessionalUser, id)
	var i ProfessionalUser
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Name,
		&i.Username,
		&i.Password,
		&i.Gender,
		&i.Email,
		&i.DateOfBirth,
		&i.Cpf,
		&i.ImageID,
		&i.PhoneID,
		&i.ProfessionalInformationID,
		&i.UpdatedAt,
		&i.SubjectMatterID,
		&i.SubjectMatterClassID,
		&i.ClassHourPrice,
		&i.CalendarID,
	)
	return i, err
}

const getProfessionalUser = `-- name: GetProfessionalUser :one
SELECT id, created_at, name, username, password, gender, email, date_of_birth, cpf, image_id, phone_id, professional_information_id, updated_at, "subjectMatter_id", "subjectMatter_class_id", class_hour_price, calendar_id FROM "professionalUser" 
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetProfessionalUser(ctx context.Context, id int64) (ProfessionalUser, error) {
	row := q.queryRow(ctx, q.getProfessionalUserStmt, getProfessionalUser, id)
	var i ProfessionalUser
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Name,
		&i.Username,
		&i.Password,
		&i.Gender,
		&i.Email,
		&i.DateOfBirth,
		&i.Cpf,
		&i.ImageID,
		&i.PhoneID,
		&i.ProfessionalInformationID,
		&i.UpdatedAt,
		&i.SubjectMatterID,
		&i.SubjectMatterClassID,
		&i.ClassHourPrice,
		&i.CalendarID,
	)
	return i, err
}

const listProfessionalUser = `-- name: ListProfessionalUser :many
SELECT id, created_at, name, username, password, gender, email, date_of_birth, cpf, image_id, phone_id, professional_information_id, updated_at, "subjectMatter_id", "subjectMatter_class_id", class_hour_price, calendar_id FROM "professionalUser" 
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListProfessionalUserParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListProfessionalUser(ctx context.Context, arg ListProfessionalUserParams) ([]ProfessionalUser, error) {
	rows, err := q.query(ctx, q.listProfessionalUserStmt, listProfessionalUser, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ProfessionalUser{}
	for rows.Next() {
		var i ProfessionalUser
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.Name,
			&i.Username,
			&i.Password,
			&i.Gender,
			&i.Email,
			&i.DateOfBirth,
			&i.Cpf,
			&i.ImageID,
			&i.PhoneID,
			&i.ProfessionalInformationID,
			&i.UpdatedAt,
			&i.SubjectMatterID,
			&i.SubjectMatterClassID,
			&i.ClassHourPrice,
			&i.CalendarID,
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

const updateProfessionalUser = `-- name: UpdateProfessionalUser :one
UPDATE "professionalUser"
    set  name = $2,
    username = $3,
    password = $4,
    email = $5, 
    date_of_birth = $6,
    class_hour_price = $7
WHERE id = $1
RETURNING id, created_at, name, username, password, gender, email, date_of_birth, cpf, image_id, phone_id, professional_information_id, updated_at, "subjectMatter_id", "subjectMatter_class_id", class_hour_price, calendar_id
`

type UpdateProfessionalUserParams struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	Username       string    `json:"username"`
	Password       string    `json:"password"`
	Email          string    `json:"email"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	ClassHourPrice string    `json:"class_hour_price"`
}

func (q *Queries) UpdateProfessionalUser(ctx context.Context, arg UpdateProfessionalUserParams) (ProfessionalUser, error) {
	row := q.queryRow(ctx, q.updateProfessionalUserStmt, updateProfessionalUser,
		arg.ID,
		arg.Name,
		arg.Username,
		arg.Password,
		arg.Email,
		arg.DateOfBirth,
		arg.ClassHourPrice,
	)
	var i ProfessionalUser
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Name,
		&i.Username,
		&i.Password,
		&i.Gender,
		&i.Email,
		&i.DateOfBirth,
		&i.Cpf,
		&i.ImageID,
		&i.PhoneID,
		&i.ProfessionalInformationID,
		&i.UpdatedAt,
		&i.SubjectMatterID,
		&i.SubjectMatterClassID,
		&i.ClassHourPrice,
		&i.CalendarID,
	)
	return i, err
}
