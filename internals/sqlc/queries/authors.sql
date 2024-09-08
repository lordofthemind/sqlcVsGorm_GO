-- name: GetAuthor :one
SELECT id, name, bio, email, date_of_birth FROM authors
WHERE id = $1;

-- name: ListAuthors :many
SELECT id, name, bio, email, date_of_birth FROM authors
ORDER BY name;

-- name: CreateAuthor :one
INSERT INTO authors (name, bio, email, date_of_birth)
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = $1;

-- name: UpdateAuthor :exec
UPDATE authors
SET name = $2,
    bio = $3,
    email = $4,
    date_of_birth = $5
WHERE id = $1;

-- name: GetAuthorsByBirthdateRange :many
SELECT id, name, bio, email, date_of_birth FROM authors
WHERE date_of_birth BETWEEN $1 AND $2
ORDER BY date_of_birth;
