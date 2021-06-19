
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



-- name: GetCreator :one
SELECT * FROM users INNER JOIN creators
ON users.fanfit_user_id = creators.fanfit_user_id
WHERE email = $1;
