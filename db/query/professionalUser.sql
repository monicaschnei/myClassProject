-- name: CreateProfessionalUser :one
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
RETURNING *;

-- name: GetProfessionalUser :one
SELECT * FROM "professionalUser" 
WHERE username = $1 LIMIT 1;

-- name: ListProfessionalUser :many
SELECT * FROM "professionalUser" 
ORDER BY id
LIMIT $1
OFFSET $2;


-- name: UpdateProfessionalUser :one
UPDATE "professionalUser"
    set  name = $2,
    username = $3,
    hashed_password = $4,
    email = $5, 
    date_of_birth = $6,
    class_hour_price = $7
WHERE id = $1
RETURNING *;

-- name: DeleteProfessionalUser :one
DELETE FROM "professionalUser"
WHERE id = $1
RETURNING *;




