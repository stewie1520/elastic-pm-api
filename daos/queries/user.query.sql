-- name: FindUserById :one
SELECT * FROM "users" WHERE "users"."id" = $1 LIMIT 1;