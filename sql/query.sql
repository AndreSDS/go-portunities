-- name: GetOpenning :one
SELECT
    *
FROM
    opennings
WHERE
    id = ?
LIMIT
    1;

-- name: ListOpennings :many
SELECT
    *
FROM
    opennings
ORDER BY
    company;

-- name: CreateOpenning :one
INSERT INTO
    opennings (role, company, location, remote, link, salary)
VALUES
    (?, ?, ?, ?, ?, ?) RETURNING *;

-- name: UpdateOpenning :exec
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

-- name: DeleteOpenning :exec
DELETE FROM opennings
WHERE
    id = ?;