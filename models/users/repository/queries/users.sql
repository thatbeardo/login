-- name: GetUserByEmail :one
SELECT * FROM users 
WHERE email = $1;


-- name: CreateUser :one
INSERT INTO users ( 
  first_name,
  last_name, 
  email,
  user_type_id,
  username,
  phone_no,
  gender,
  profile_picture,
  bio,
  background_picture
) VALUES (
  $1, 
  $2, 
  $3,
  $4, 
  $5, 
  $6,
  $7, 
  $8, 
  $9,
  $10
)
RETURNING *;

-- name: CreateClient :one
INSERT INTO clients(
  fanfit_user_id,
  temp_field
) VALUES(
  $1,
  $2
)
RETURNING *;


-- name: DeleteUser :one
DELETE FROM users
WHERE email = $1
RETURNING *;

-- name: GetClientByID :one
SELECT * FROM users INNER JOIN clients
ON users.id = clients.fanfit_user_id
WHERE fanfit_user_id = $1;
