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
    hashed_password,
    gender,
    email, 
    date_of_birth,
    cpf,
    image_id,
    updated_at,
    class_hour_price
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
RETURNING id, created_at, name, username, gender, email, date_of_birth, cpf, image_id, password_changed_at, hashed_password, updated_at, class_hour_price
`

type CreateProfessionalUserParams struct {
	Name           string    `json:"name"`
	Username       string    `json:"username"`
	HashedPassword string    `json:"hashed_password"`
	Gender         string    `json:"gender"`
	Email          string    `json:"email"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	Cpf            string    `json:"cpf"`
	ImageID        int64     `json:"image_id"`
	UpdatedAt      time.Time `json:"updated_at"`
	ClassHourPrice string    `json:"class_hour_price"`
}

func (q *Queries) CreateProfessionalUser(ctx context.Context, arg CreateProfessionalUserParams) (ProfessionalUser, error) {
	row := q.queryRow(ctx, q.createProfessionalUserStmt, createProfessionalUser,
		arg.Name,
		arg.Username,
		arg.HashedPassword,
		arg.Gender,
		arg.Email,
		arg.DateOfBirth,
		arg.Cpf,
		arg.ImageID,
		arg.UpdatedAt,
		arg.ClassHourPrice,
	)
	var i ProfessionalUser
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Name,
		&i.Username,
		&i.Gender,
		&i.Email,
		&i.DateOfBirth,
		&i.Cpf,
		&i.ImageID,
		&i.PasswordChangedAt,
		&i.HashedPassword,
		&i.UpdatedAt,
		&i.ClassHourPrice,
	)
	return i, err
}

const deleteProfessionalUser = `-- name: DeleteProfessionalUser :one
DELETE FROM "professionalUser"
WHERE id = $1
RETURNING id, created_at, name, username, gender, email, date_of_birth, cpf, image_id, password_changed_at, hashed_password, updated_at, class_hour_price
`

func (q *Queries) DeleteProfessionalUser(ctx context.Context, id int64) (ProfessionalUser, error) {
	row := q.queryRow(ctx, q.deleteProfessionalUserStmt, deleteProfessionalUser, id)
	var i ProfessionalUser
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Name,
		&i.Username,
		&i.Gender,
		&i.Email,
		&i.DateOfBirth,
		&i.Cpf,
		&i.ImageID,
		&i.PasswordChangedAt,
		&i.HashedPassword,
		&i.UpdatedAt,
		&i.ClassHourPrice,
	)
	return i, err
}

const getProfessionalUser = `-- name: GetProfessionalUser :one
SELECT id, created_at, name, username, gender, email, date_of_birth, cpf, image_id, password_changed_at, hashed_password, updated_at, class_hour_price FROM "professionalUser" 
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetProfessionalUser(ctx context.Context, username string) (ProfessionalUser, error) {
	row := q.queryRow(ctx, q.getProfessionalUserStmt, getProfessionalUser, username)
	var i ProfessionalUser
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Name,
		&i.Username,
		&i.Gender,
		&i.Email,
		&i.DateOfBirth,
		&i.Cpf,
		&i.ImageID,
		&i.PasswordChangedAt,
		&i.HashedPassword,
		&i.UpdatedAt,
		&i.ClassHourPrice,
	)
	return i, err
}

const listProfessionalUser = `-- name: ListProfessionalUser :many
SELECT id, created_at, name, username, gender, email, date_of_birth, cpf, image_id, password_changed_at, hashed_password, updated_at, class_hour_price FROM "professionalUser" 
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
			&i.Gender,
			&i.Email,
			&i.DateOfBirth,
			&i.Cpf,
			&i.ImageID,
			&i.PasswordChangedAt,
			&i.HashedPassword,
			&i.UpdatedAt,
			&i.ClassHourPrice,
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
    hashed_password = $4,
    email = $5, 
    date_of_birth = $6,
    class_hour_price = $7
WHERE id = $1
RETURNING id, created_at, name, username, gender, email, date_of_birth, cpf, image_id, password_changed_at, hashed_password, updated_at, class_hour_price
`

type UpdateProfessionalUserParams struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	Username       string    `json:"username"`
	HashedPassword string    `json:"hashed_password"`
	Email          string    `json:"email"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	ClassHourPrice string    `json:"class_hour_price"`
}

func (q *Queries) UpdateProfessionalUser(ctx context.Context, arg UpdateProfessionalUserParams) (ProfessionalUser, error) {
	row := q.queryRow(ctx, q.updateProfessionalUserStmt, updateProfessionalUser,
		arg.ID,
		arg.Name,
		arg.Username,
		arg.HashedPassword,
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
		&i.Gender,
		&i.Email,
		&i.DateOfBirth,
		&i.Cpf,
		&i.ImageID,
		&i.PasswordChangedAt,
		&i.HashedPassword,
		&i.UpdatedAt,
		&i.ClassHourPrice,
	)
	return i, err
}
