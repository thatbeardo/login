
-- name: CreateCreator :one
INSERT INTO creators(
  fanfit_user_id,
  payment_info,
  logo_picture
) VALUES(
  $1,
  $2,
  $3
)
RETURNING *;



-- name: GetCreatorByEmail :one
SELECT * FROM users INNER JOIN creators
ON users.id = creators.fanfit_user_id
WHERE email = $1;
