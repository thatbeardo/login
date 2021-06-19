package repository

import (
	"context"
	"fmt"
)

// Repository is used by the service to communicate with the underlying database
type Repository interface {
	GetCreator(context.Context, string) (GetCreatorRow, error)
	CreateCreator(context.Context, Creator) (Creator, error)
}

type repository struct {
	queries *Queries
}

// GetCreator with fan_fit_userid
func (repo *repository) GetCreator(ctx context.Context, FanfitUserID string) (GetCreatorRow, error) {
	temp, err := repo.queries.GetCreator(ctx, FanfitUserID)

	if err != nil {
		fmt.Print(err)
	}

	return temp, err

}

// Create Users
func (repo *repository) CreateCreator(ctx context.Context, cons Creator) (Creator, error) {
	response, err := repo.queries.CreateCreator(ctx, CreateCreatorParams{
		FanfitUserID:      cons.FanfitUserID,
		PaymentInfo:       cons.PaymentInfo,
		LogoPicture:       cons.LogoPicture,
		BackgroundPicture: cons.BackgroundPicture,
	})

	if err != nil {
		fmt.Print(err)
	}
	return response, err
}

func NewUserStore(db DBTX) Repository {
	return &repository{
		queries: New(db),
	}
}
