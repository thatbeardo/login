-- name: GetUserByEmail :one
SELECT * FROM users 
WHERE email = $1;

-- name: GetUserByUsername :one
SELECT * FROM users 
WHERE username = $1;

-- name: CreateUser :one
INSERT INTO users ( 
  first_name,
  last_name, 
  email,
  created_date,
  username,
  phone_no,
  gender,
  profile_picture,
  bio
) VALUES (
  $1, 
  $2, 
  $3,
  $4, 
  $5, 
  $6,
  $7, 
  $8, 
  $9
)
RETURNING *;

-- name: DeleteUser :one
DELETE FROM users
WHERE email = $1
RETURNING *;