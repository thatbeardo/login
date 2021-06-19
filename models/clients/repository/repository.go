package repository

import (
	"context"
	"fmt"
)

// Repository is used by the service to communicate with the underlying database
type Repository interface {
	Delete(context.Context, string) error
	CreateClient(context.Context, Client) (Client, error)
	GetClientByEmail(context.Context, string) (GetClientByEmailRow, error)
}

type repository struct {
	queries *Queries
}

// GetClientByEmail
func (repo *repository) GetClientByEmail(ctx context.Context, FanfitUserID string) (GetClientByEmailRow, error) {
	response, err := repo.queries.GetClientByEmail(ctx, FanfitUserID)
	if err != nil {
		fmt.Print(err)
	}
	return response, err
}

// Delete function deletes a node from the graph
func (repo *repository) Delete(ctx context.Context, id string) error {
	return nil
}

func NewUserStore(db DBTX) Repository {
	return &repository{
		queries: New(db),
	}
}

// Create Users
func (repo *repository) CreateClient(ctx context.Context, cons Client) (Client, error) {
	response, err := repo.queries.CreateClient(ctx, CreateClientParams{
		FanfitUserID: cons.FanfitUserID,
		TempField:    cons.TempField,
	})

	if err != nil {
		fmt.Print(err)
	}
	return response, err
}
