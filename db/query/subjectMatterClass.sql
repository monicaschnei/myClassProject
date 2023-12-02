-- name: CreateSubjectMatterClass :one
INSERT INTO "subjectMatterClass" (
"subjectMatter_id",
professional_id,
durantion,
enrollment_date,
enrollment_time,
cancellation,
cancellation_reason,
student_attendence,
study_material,
testing_exam
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
RETURNING *;

-- name: GetSubjectMatterClass :one
SELECT * FROM "subjectMatterClass" 
WHERE id = $1 LIMIT 1;

-- name: ListSubjectMatterClass :many
SELECT * FROM "subjectMatterClass" 
ORDER BY id
LIMIT $1
OFFSET $2;


-- name: UpdateSubjectMatterClass :one
UPDATE "subjectMatterClass"
    set durantion = $2,
    enrollment_date = $3,
    enrollment_time = $4,
    study_material = $5,
    testing_exam = $6
WHERE id = $1
RETURNING *;

-- name: DeleteSubjectMatterClass :one
DELETE FROM "subjectMatterClass"
WHERE id = $1
RETURNING *;
