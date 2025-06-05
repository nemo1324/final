-- name: CreateUser :one
INSERT INTO users (name, login, pass, status)
VALUES (@name, @login, @pass, @status)
RETURNING id::BIGINT;

-- name: GetUserByLogin :one
SELECT * FROM users
WHERE login = @login;


-- name: GetUserByID :one
SELECT * FROM users
WHERE id = @id;

