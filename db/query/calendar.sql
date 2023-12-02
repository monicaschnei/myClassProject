-- name: CreateCalendar :one
INSERT INTO "calendar" (
  "subjectMatter_id",
  time,
  date,
  available,
  filled_student_id,
  "professionalUser_id"
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetCalendar :one
SELECT * FROM  "calendar"
WHERE id = $1 LIMIT 1;

-- name: ListCalendar :many
SELECT * FROM  "calendar"
ORDER BY id
LIMIT $1
OFFSET $2;


-- name: UpdateCalendar :one
UPDATE  "calendar"
    set  time = $2,
    date = $3,
    available = $4
WHERE id = $1
RETURNING *;

-- name: DeleteCalendar :one
DELETE FROM  "calendar"
WHERE id = $1
RETURNING *;

