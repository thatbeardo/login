-- name: GetUserByEmail :one
SELECT * FROM users 
WHERE email = $1;

-- name: CreateUser :one
INSERT INTO users ( 
  first_name,
  last_name, 
  email
) VALUES (
  $1, 
  $2, 
  $3
)
RETURNING *;

-- name: DeleteUser :one
DELETE FROM users
WHERE email = $1
RETURNING *;