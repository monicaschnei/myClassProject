-- name: CreateResponsibleStudent :one
INSERT INTO "responsibleStudent" (
name,
gender,
email,
date_of_birth,
username,
cpf,
hashed_password,
updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: GetResponsibleStudent :one
SELECT * FROM "responsibleStudent" 
WHERE username = $1 LIMIT 1;

-- name: ListResponsibleStudent :many
SELECT * FROM "responsibleStudent" 
ORDER BY id
LIMIT $1
OFFSET $2;


-- name: UpdateResponsibleStudent :one
UPDATE "responsibleStudent"
    set name = $2,
    email = $3
WHERE id = $1
RETURNING *;

-- name: DeleteResponsibleStudent :one
DELETE FROM "responsibleStudent"
WHERE id = $1
RETURNING *;