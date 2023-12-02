-- name: CreateStudentUser :one
INSERT INTO "studentUser" (
    username,
    password,
    name,
    date_of_birth,
    gender,
    responsible_student_id,
    updated_at,
    "subjectMatter_class_id",
    calendar_id
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- name: GetStudentUser :one
SELECT * FROM "studentUser" 
WHERE id = $1 LIMIT 1;

-- name: ListStudentUser :many
SELECT * FROM "studentUser" 
ORDER BY id
LIMIT $1
OFFSET $2;


-- name: UpdateStudentUser :one
UPDATE "studentUser"
    set username = $2,
    password = $3,
    name = $4,
    responsible_student_id = $5,
    "subjectMatter_class_id"  = $6
WHERE id = $1
RETURNING *;

-- name: DeleteStudentUser :one
DELETE FROM "studentUser"
WHERE id = $1
RETURNING *;
