// Code generated by sqlc. DO NOT EDIT.
// source: creators.sql

package repository

import (
	"context"
	"database/sql"
	"time"
)

const createCreator = `-- name: CreateCreator :one
INSERT INTO creators(
  fanfit_user_id,
  payment_info,
  logo_picture
) VALUES(
  $1,
  $2,
  $3
)
RETURNING fanfit_user_id, payment_info, logo_picture
`

type CreateCreatorParams struct {
	FanfitUserID int32
	PaymentInfo  string
	LogoPicture  string
}

func (q *Queries) CreateCreator(ctx context.Context, arg CreateCreatorParams) (Creator, error) {
	row := q.db.QueryRowContext(ctx, createCreator, arg.FanfitUserID, arg.PaymentInfo, arg.LogoPicture)
	var i Creator
	err := row.Scan(&i.FanfitUserID, &i.PaymentInfo, &i.LogoPicture)
	return i, err
}

const getCreatorByEmail = `-- name: GetCreatorByEmail :one
SELECT id, user_type_id, first_name, last_name, email, created_date, username, phone_no, gender, profile_picture, bio, background_picture, fanfit_user_id, payment_info, logo_picture FROM users INNER JOIN creators
ON users.id = creators.fanfit_user_id
WHERE email = $1
`

type GetCreatorByEmailRow struct {
	ID                int32
	UserTypeID        int32
	FirstName         string
	LastName          string
	Email             string
	CreatedDate       time.Time
	Username          sql.NullString
	PhoneNo           sql.NullString
	Gender            sql.NullString
	ProfilePicture    sql.NullString
	Bio               sql.NullString
	BackgroundPicture sql.NullString
	FanfitUserID      int32
	PaymentInfo       string
	LogoPicture       string
}

func (q *Queries) GetCreatorByEmail(ctx context.Context, email string) (GetCreatorByEmailRow, error) {
	row := q.db.QueryRowContext(ctx, getCreatorByEmail, email)
	var i GetCreatorByEmailRow
	err := row.Scan(
		&i.ID,
		&i.UserTypeID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.CreatedDate,
		&i.Username,
		&i.PhoneNo,
		&i.Gender,
		&i.ProfilePicture,
		&i.Bio,
		&i.BackgroundPicture,
		&i.FanfitUserID,
		&i.PaymentInfo,
		&i.LogoPicture,
	)
	return i, err
}