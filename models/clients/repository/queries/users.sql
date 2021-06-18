
-- name: CreateClients :one
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


-- name: GetClients :one
SELECT * FROM users FULL OUTER JOIN clients
ON users.fanfit_user_id = clients.fanfit_user_id
WHERE fanfit_user_id = $1;
