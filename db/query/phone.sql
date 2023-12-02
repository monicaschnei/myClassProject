-- name: CreatePhone :one
INSERT INTO "phone" (
  country_code,
  area_core,
  number,
  type,
  updated_at
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetPhone :one
SELECT * FROM "phone" 
WHERE id = $1 LIMIT 1;

-- name: ListPhone :many
SELECT * FROM "phone" 
ORDER BY id
LIMIT $1
OFFSET $2;


-- name: UpdatePhone :one
UPDATE "phone"
    set country_code = $2,
    area_core = $3,
    number = $4,
    type = $5
WHERE id = $1
RETURNING *;

-- name: DeletePhone :one
DELETE FROM "phone"
WHERE id = $1
RETURNING *;