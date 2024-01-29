// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: studentUser.sql

package db

import (
	"context"
	"time"
)

const createStudentUser = `-- name: CreateStudentUser :one
INSERT INTO "studentUser" (
    username,
    password,
    name,
    date_of_birth,
    gender,
    responsible_student_id,
    updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING id, username, password, name, date_of_birth, gender, created_at, responsible_student_id, updated_at
`

type CreateStudentUserParams struct {
	Username             string    `json:"username"`
	Password             string    `json:"password"`
	Name                 string    `json:"name"`
	DateOfBirth          time.Time `json:"date_of_birth"`
	Gender               string    `json:"gender"`
	ResponsibleStudentID int32     `json:"responsible_student_id"`
	UpdatedAt            time.Time `json:"updated_at"`
}

func (q *Queries) CreateStudentUser(ctx context.Context, arg CreateStudentUserParams) (StudentUser, error) {
	row := q.queryRow(ctx, q.createStudentUserStmt, createStudentUser,
		arg.Username,
		arg.Password,
		arg.Name,
		arg.DateOfBirth,
		arg.Gender,
		arg.ResponsibleStudentID,
		arg.UpdatedAt,
	)
	var i StudentUser
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Name,
		&i.DateOfBirth,
		&i.Gender,
		&i.CreatedAt,
		&i.ResponsibleStudentID,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteStudentUser = `-- name: DeleteStudentUser :one
DELETE FROM "studentUser"
WHERE id = $1
RETURNING id, username, password, name, date_of_birth, gender, created_at, responsible_student_id, updated_at
`

func (q *Queries) DeleteStudentUser(ctx context.Context, id int64) (StudentUser, error) {
	row := q.queryRow(ctx, q.deleteStudentUserStmt, deleteStudentUser, id)
	var i StudentUser
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Name,
		&i.DateOfBirth,
		&i.Gender,
		&i.CreatedAt,
		&i.ResponsibleStudentID,
		&i.UpdatedAt,
	)
	return i, err
}

const getStudentUser = `-- name: GetStudentUser :one
SELECT id, username, password, name, date_of_birth, gender, created_at, responsible_student_id, updated_at FROM "studentUser" 
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetStudentUser(ctx context.Context, id int64) (StudentUser, error) {
	row := q.queryRow(ctx, q.getStudentUserStmt, getStudentUser, id)
	var i StudentUser
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Name,
		&i.DateOfBirth,
		&i.Gender,
		&i.CreatedAt,
		&i.ResponsibleStudentID,
		&i.UpdatedAt,
	)
	return i, err
}

const listStudentUser = `-- name: ListStudentUser :many
SELECT id, username, password, name, date_of_birth, gender, created_at, responsible_student_id, updated_at FROM "studentUser" 
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListStudentUserParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListStudentUser(ctx context.Context, arg ListStudentUserParams) ([]StudentUser, error) {
	rows, err := q.query(ctx, q.listStudentUserStmt, listStudentUser, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []StudentUser{}
	for rows.Next() {
		var i StudentUser
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Password,
			&i.Name,
			&i.DateOfBirth,
			&i.Gender,
			&i.CreatedAt,
			&i.ResponsibleStudentID,
			&i.UpdatedAt,
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

const updateStudentUser = `-- name: UpdateStudentUser :one
UPDATE "studentUser"
    set username = $2,
    password = $3,
    name = $4,
    responsible_student_id = $5
WHERE id = $1
RETURNING id, username, password, name, date_of_birth, gender, created_at, responsible_student_id, updated_at
`

type UpdateStudentUserParams struct {
	ID                   int64  `json:"id"`
	Username             string `json:"username"`
	Password             string `json:"password"`
	Name                 string `json:"name"`
	ResponsibleStudentID int32  `json:"responsible_student_id"`
}

func (q *Queries) UpdateStudentUser(ctx context.Context, arg UpdateStudentUserParams) (StudentUser, error) {
	row := q.queryRow(ctx, q.updateStudentUserStmt, updateStudentUser,
		arg.ID,
		arg.Username,
		arg.Password,
		arg.Name,
		arg.ResponsibleStudentID,
	)
	var i StudentUser
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Name,
		&i.DateOfBirth,
		&i.Gender,
		&i.CreatedAt,
		&i.ResponsibleStudentID,
		&i.UpdatedAt,
	)
	return i, err
}
