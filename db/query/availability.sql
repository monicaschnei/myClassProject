-- name: AddAvailability :one
INSERT INTO "availability" (
  date,
  start,
  end_time,
  is_available,
  user_id,
  username
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetAvailability :one
SELECT * FROM  "availability"
WHERE id = $1 LIMIT 1;

-- name: ListAvailability :many
SELECT * FROM  "availability"
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteAvailability :one
DELETE FROM  "availability"
WHERE id = $1
RETURNING *;
