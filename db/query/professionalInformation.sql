-- name: CreateProfessionalInformation :one
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
RETURNING *;

-- name: GetProfessionalInformation :one
SELECT * FROM "professionalInformation" 
WHERE id = $1 LIMIT 1;

-- name: ListProfessionalInformation :many
SELECT * FROM "professionalInformation" 
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: ListProfessionalInformationByUser :many
SELECT * FROM "professionalInformation"
WHERE professional_user_id = $1
LIMIT $1
OFFSET $2;

-- name: UpdateProfessionalInformation :one
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
RETURNING *;

-- name: DeleteProfessionalInformation :one
DELETE FROM "professionalInformation"
WHERE id = $1
RETURNING *;