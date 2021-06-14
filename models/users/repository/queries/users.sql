-- name: GetUserByEmail :one
SELECT * FROM users 
WHERE email = $1;

-- name: GetClient :one
SELECT * FROM users
INNER JOIN consumers on consumers.fanfit_user_id = users.id
WHERE users.email = $1;

-- name: GetCreator :one
SELECT * FROM users
INNER JOIN creators on creators.fanfit_user_id = users.id
WHERE users.email = $1;

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

-- name: CreateConsumer :one
INSERT INTO consumers(
  fanfit_user_id,
  temp_field
) VALUES(
  $1,
  $2
)
RETURNING *;

-- name: CreateCreator :one
INSERT INTO creators(
  fanfit_user_id,
  payment_info,
  logo_picture,
  background_picture
) VALUES(
  $1,
  $2,
  $3,
  $4
)
RETURNING *;


-- name: DeleteUser :one
DELETE FROM users
WHERE email = $1
RETURNING *;