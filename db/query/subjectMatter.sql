-- name: CreateSubjectMatter :one
INSERT INTO "subjectMatter" (
    title,
    category,
    abstract
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetSubjectMatter :one
SELECT * FROM "subjectMatter" 
WHERE id = $1 LIMIT 1;

-- name: ListSubjectMatter :many
SELECT * FROM "subjectMatter" 
ORDER BY id
LIMIT $1
OFFSET $2;


-- name: UpdateSubjectMatter :one
UPDATE "subjectMatter"
    set title = $2,
    category = $3,
    abstract = $4
WHERE id = $1
RETURNING *;

-- name: DeleteSubjectMatter :one
DELETE FROM "subjectMatter"
WHERE id = $1
RETURNING *;
