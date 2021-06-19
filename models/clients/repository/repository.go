package repository

import (
	"context"
	"database/sql"
	"fmt"
)

// Repository is used by the service to communicate with the underlying database
type Repository interface {
	Delete(context.Context, string) error
	CreateClient(context.Context, int32) (GetClientByIDRow, error)
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

// Create Client
func (repo *repository) CreateClient(ctx context.Context, userID int32) (GetClientByIDRow, error) {
	response, err := repo.queries.CreateClient(ctx, CreateClientParams{
		FanfitUserID: userID,
		TempField:    sql.NullString{},
	})
	if err != nil {
		fmt.Print(err)
	}
	client, err2 := repo.queries.GetClientByID(ctx, response.FanfitUserID)
	if err2 != nil {
		fmt.Print(err2)
	}
	return client, err2
}
