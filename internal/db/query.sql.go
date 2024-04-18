// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package db

import (
	"context"
)

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO
    opennings (role, company, location, remote, link, salary)
VALUES
    (?, ?, ?, ?, ?, ?) RETURNING id, role, company, location, remote, link, salary
`

type CreateAuthorParams struct {
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int64
}

func (q *Queries) CreateAuthor(ctx context.Context, arg CreateAuthorParams) (Openning, error) {
	row := q.db.QueryRowContext(ctx, createAuthor,
		arg.Role,
		arg.Company,
		arg.Location,
		arg.Remote,
		arg.Link,
		arg.Salary,
	)
	var i Openning
	err := row.Scan(
		&i.ID,
		&i.Role,
		&i.Company,
		&i.Location,
		&i.Remote,
		&i.Link,
		&i.Salary,
	)
	return i, err
}

const deleteAuthor = `-- name: DeleteAuthor :exec
DELETE FROM opennings
WHERE
    id = ?
`

func (q *Queries) DeleteAuthor(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAuthor, id)
	return err
}

const getAuthor = `-- name: GetAuthor :one
SELECT
    id, role, company, location, remote, link, salary
FROM
    opennings
WHERE
    id = ?
LIMIT
    1
`

func (q *Queries) GetAuthor(ctx context.Context, id int64) (Openning, error) {
	row := q.db.QueryRowContext(ctx, getAuthor, id)
	var i Openning
	err := row.Scan(
		&i.ID,
		&i.Role,
		&i.Company,
		&i.Location,
		&i.Remote,
		&i.Link,
		&i.Salary,
	)
	return i, err
}

const listAuthors = `-- name: ListAuthors :many
SELECT
    id, role, company, location, remote, link, salary
FROM
    opennings
ORDER BY
    company
`

func (q *Queries) ListAuthors(ctx context.Context) ([]Openning, error) {
	rows, err := q.db.QueryContext(ctx, listAuthors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Openning
	for rows.Next() {
		var i Openning
		if err := rows.Scan(
			&i.ID,
			&i.Role,
			&i.Company,
			&i.Location,
			&i.Remote,
			&i.Link,
			&i.Salary,
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

const updateAuthor = `-- name: UpdateAuthor :exec
UPDATE opennings
SET
    role = ?,
    company = ?,
    location = ?,
    remote = ?,
    link = ?,
    salary = ?
WHERE
    id = ?
`

type UpdateAuthorParams struct {
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int64
	ID       int64
}

func (q *Queries) UpdateAuthor(ctx context.Context, arg UpdateAuthorParams) error {
	_, err := q.db.ExecContext(ctx, updateAuthor,
		arg.Role,
		arg.Company,
		arg.Location,
		arg.Remote,
		arg.Link,
		arg.Salary,
		arg.ID,
	)
	return err
}