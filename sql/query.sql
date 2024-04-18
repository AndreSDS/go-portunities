-- name: GetAuthor :one
SELECT
    *
FROM
    opennings
WHERE
    id = ?
LIMIT
    1;

-- name: ListAuthors :many
SELECT
    *
FROM
    opennings
ORDER BY
    company;

-- name: CreateAuthor :one
INSERT INTO
    opennings (role, company, location, remote, link, salary)
VALUES
    (?, ?, ?, ?, ?, ?) RETURNING *;

-- name: UpdateAuthor :exec
UPDATE opennings
SET
    role = ?,
    company = ?,
    location = ?,
    remote = ?,
    link = ?,
    salary = ?
WHERE
    id = ?;

-- name: DeleteAuthor :exec
DELETE FROM opennings
WHERE
    id = ?;