// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: phone.sql

package db

import (
	"context"
	"time"
)

const createPhone = `-- name: CreatePhone :one
INSERT INTO "phone" (
  country_code,
  area_core,
  number,
  type,
  updated_at
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING id, country_code, area_core, number, type, created_at, updated_at, user_id
`

type CreatePhoneParams struct {
	CountryCode int32     `json:"country_code"`
	AreaCore    int32     `json:"area_core"`
	Number      int32     `json:"number"`
	Type        string    `json:"type"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (q *Queries) CreatePhone(ctx context.Context, arg CreatePhoneParams) (Phone, error) {
	row := q.queryRow(ctx, q.createPhoneStmt, createPhone,
		arg.CountryCode,
		arg.AreaCore,
		arg.Number,
		arg.Type,
		arg.UpdatedAt,
	)
	var i Phone
	err := row.Scan(
		&i.ID,
		&i.CountryCode,
		&i.AreaCore,
		&i.Number,
		&i.Type,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
	)
	return i, err
}

const deletePhone = `-- name: DeletePhone :one
DELETE FROM "phone"
WHERE id = $1
RETURNING id, country_code, area_core, number, type, created_at, updated_at, user_id
`

func (q *Queries) DeletePhone(ctx context.Context, id int64) (Phone, error) {
	row := q.queryRow(ctx, q.deletePhoneStmt, deletePhone, id)
	var i Phone
	err := row.Scan(
		&i.ID,
		&i.CountryCode,
		&i.AreaCore,
		&i.Number,
		&i.Type,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
	)
	return i, err
}

const getPhone = `-- name: GetPhone :one
SELECT id, country_code, area_core, number, type, created_at, updated_at, user_id FROM "phone" 
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPhone(ctx context.Context, id int64) (Phone, error) {
	row := q.queryRow(ctx, q.getPhoneStmt, getPhone, id)
	var i Phone
	err := row.Scan(
		&i.ID,
		&i.CountryCode,
		&i.AreaCore,
		&i.Number,
		&i.Type,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
	)
	return i, err
}

const listPhone = `-- name: ListPhone :many
SELECT id, country_code, area_core, number, type, created_at, updated_at, user_id FROM "phone" 
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListPhoneParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListPhone(ctx context.Context, arg ListPhoneParams) ([]Phone, error) {
	rows, err := q.query(ctx, q.listPhoneStmt, listPhone, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Phone{}
	for rows.Next() {
		var i Phone
		if err := rows.Scan(
			&i.ID,
			&i.CountryCode,
			&i.AreaCore,
			&i.Number,
			&i.Type,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
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

const updatePhone = `-- name: UpdatePhone :one
UPDATE "phone"
    set country_code = $2,
    area_core = $3,
    number = $4,
    type = $5
WHERE id = $1
RETURNING id, country_code, area_core, number, type, created_at, updated_at, user_id
`

type UpdatePhoneParams struct {
	ID          int64  `json:"id"`
	CountryCode int32  `json:"country_code"`
	AreaCore    int32  `json:"area_core"`
	Number      int32  `json:"number"`
	Type        string `json:"type"`
}

func (q *Queries) UpdatePhone(ctx context.Context, arg UpdatePhoneParams) (Phone, error) {
	row := q.queryRow(ctx, q.updatePhoneStmt, updatePhone,
		arg.ID,
		arg.CountryCode,
		arg.AreaCore,
		arg.Number,
		arg.Type,
	)
	var i Phone
	err := row.Scan(
		&i.ID,
		&i.CountryCode,
		&i.AreaCore,
		&i.Number,
		&i.Type,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
	)
	return i, err
}
