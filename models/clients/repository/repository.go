package repository

import (
	"context"
	"fmt"
)

// Repository is used by the service to communicate with the underlying database
type Repository interface {
	Delete(context.Context, string) error
	CreateClients(context.Context, Client) (Client, error)
	GetClients(context.Context, int32) (GetClientsRow, error)
}

type repository struct {
	queries *Queries
}

// GetClients
func (repo *repository) GetClients(ctx context.Context, FanfitUserID int32) (GetClientsRow, error) {
	response, err := repo.queries.GetClients(ctx, FanfitUserID)
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
func (repo *repository) CreateClients(ctx context.Context, cons Client) (Client, error) {
	response, err := repo.queries.CreateClients(ctx, CreateClientsParams{
		FanfitUserID: cons.FanfitUserID,
		TempField:    cons.TempField,
	})

	if err != nil {
		fmt.Print(err)
	}
	return response, err
}
