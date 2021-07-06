package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/fanfit/user-service/database"
)

// Repository is used by the service to communicate with the underlying database
type Repository interface {
	Delete(context.Context, string) error
	CreateClient(context.Context, Client) (GetClientByIDRow, error)
	GetClientByEmail(context.Context, string) (GetClientByEmailRow, error)
	Close() error
}

type repository struct {
	queries *Queries
	db      *sql.DB
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

// Create Client
func (repo *repository) CreateClient(ctx context.Context, client Client) (GetClientByIDRow, error) {
	response, err := repo.queries.CreateClient(ctx, CreateClientParams(client))
	if err != nil {
		fmt.Print(err)
	}
	fullClientObj, err2 := repo.queries.GetClientByID(ctx, response.FanfitUserID)
	if err2 != nil {
		fmt.Print(err2)
	}
	return fullClientObj, err2
}

func (repo *repository) Close() error {
	return repo.db.Close()
}

func NewClientStore(dbURL string) (Repository, error) {
	db, err := database.EstablishConnection(dbURL)
	if err != nil {
		fmt.Println("Error while establishing connection with databse " + err.Error())
	}

	return &repository{
		db:      db,
		queries: New(db),
	}, nil
}
