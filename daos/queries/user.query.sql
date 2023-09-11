-- name: FindUserById :one
SELECT * FROM "users" WHERE "users"."id" = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO "users" ("id", "fullName", "accountId", "createdAt", "updatedAt") VALUES ($1, $2, $3, $4, $5) RETURNING *;