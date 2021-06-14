// Code generated by sqlc. DO NOT EDIT.

package repository

import (
	"context"
)

type Querier interface {
	CreateConsumer(ctx context.Context, arg CreateConsumerParams) (Consumer, error)
	CreateCreator(ctx context.Context, arg CreateCreatorParams) (Creator, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteUser(ctx context.Context, email string) (User, error)
	GetConsumer(ctx context.Context, fanfitUserID int32) (Consumer, error)
	GetCreator(ctx context.Context, email string) (GetCreatorRow, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
}

var _ Querier = (*Queries)(nil)
