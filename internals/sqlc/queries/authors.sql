-- name: GetAuthor :one
SELECT id, name, bio FROM authors
WHERE id = $1;  -- Use $1 for PostgreSQL parameter placeholders

-- name: ListAuthors :many
SELECT id, name, bio FROM authors
ORDER BY name;

-- name: CreateAuthor :execresult
INSERT INTO authors (name, bio)
VALUES ($1, $2)
RETURNING id;  -- Return the newly created ID for the inserted author

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = $1;
