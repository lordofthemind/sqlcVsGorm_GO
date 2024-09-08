// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: authors.sql

package sqlcgen

import (
	"context"
	"database/sql"
)

const CreateAuthor = `-- name: CreateAuthor :one
INSERT INTO authors (name, bio, email, date_of_birth)
VALUES ($1, $2, $3, $4)
RETURNING id
`

type CreateAuthorParams struct {
	Name        string         `json:"name"`
	Bio         sql.NullString `json:"bio"`
	Email       string         `json:"email"`
	DateOfBirth sql.NullTime   `json:"date_of_birth"`
}

func (q *Queries) CreateAuthor(ctx context.Context, arg CreateAuthorParams) (int32, error) {
	row := q.queryRow(ctx, q.createAuthorStmt, CreateAuthor,
		arg.Name,
		arg.Bio,
		arg.Email,
		arg.DateOfBirth,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const DeleteAuthor = `-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = $1
`

func (q *Queries) DeleteAuthor(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteAuthorStmt, DeleteAuthor, id)
	return err
}

const GetAuthor = `-- name: GetAuthor :one
SELECT id, name, bio, email, date_of_birth FROM authors
WHERE id = $1
`

func (q *Queries) GetAuthor(ctx context.Context, id int32) (Author, error) {
	row := q.queryRow(ctx, q.getAuthorStmt, GetAuthor, id)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Bio,
		&i.Email,
		&i.DateOfBirth,
	)
	return i, err
}

const GetAuthorsByBirthdateRange = `-- name: GetAuthorsByBirthdateRange :many
SELECT id, name, bio, email, date_of_birth FROM authors
WHERE date_of_birth BETWEEN $1 AND $2
ORDER BY date_of_birth
`

type GetAuthorsByBirthdateRangeParams struct {
	DateOfBirth   sql.NullTime `json:"date_of_birth"`
	DateOfBirth_2 sql.NullTime `json:"date_of_birth_2"`
}

func (q *Queries) GetAuthorsByBirthdateRange(ctx context.Context, arg GetAuthorsByBirthdateRangeParams) ([]Author, error) {
	rows, err := q.query(ctx, q.getAuthorsByBirthdateRangeStmt, GetAuthorsByBirthdateRange, arg.DateOfBirth, arg.DateOfBirth_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Author{}
	for rows.Next() {
		var i Author
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Bio,
			&i.Email,
			&i.DateOfBirth,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const ListAuthors = `-- name: ListAuthors :many
SELECT id, name, bio, email, date_of_birth FROM authors
ORDER BY name
`

func (q *Queries) ListAuthors(ctx context.Context) ([]Author, error) {
	rows, err := q.query(ctx, q.listAuthorsStmt, ListAuthors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Author{}
	for rows.Next() {
		var i Author
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Bio,
			&i.Email,
			&i.DateOfBirth,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const UpdateAuthor = `-- name: UpdateAuthor :exec
UPDATE authors
SET name = $2,
    bio = $3,
    email = $4,
    date_of_birth = $5
WHERE id = $1
`

type UpdateAuthorParams struct {
	ID          int32          `json:"id"`
	Name        string         `json:"name"`
	Bio         sql.NullString `json:"bio"`
	Email       string         `json:"email"`
	DateOfBirth sql.NullTime   `json:"date_of_birth"`
}

func (q *Queries) UpdateAuthor(ctx context.Context, arg UpdateAuthorParams) error {
	_, err := q.exec(ctx, q.updateAuthorStmt, UpdateAuthor,
		arg.ID,
		arg.Name,
		arg.Bio,
		arg.Email,
		arg.DateOfBirth,
	)
	return err
}
