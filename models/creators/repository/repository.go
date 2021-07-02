package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/fanfit/login/database"
)

// Repository is used by the service to communicate with the underlying database
type Repository interface {
	GetCreatorByEmail(context.Context, string) (GetCreatorByEmailRow, error)
	CreateCreator(context.Context, Creator) (Creator, error)
	Close() error
}

type repository struct {
	queries *Queries
	db      *sql.DB
}

// GetCreatorByEmail with fan_fit_userid
func (repo *repository) GetCreatorByEmail(ctx context.Context, FanfitUserID string) (GetCreatorByEmailRow, error) {
	temp, err := repo.queries.GetCreatorByEmail(ctx, FanfitUserID)

	if err != nil {
		fmt.Print(err)
	}

	return temp, err
}

// Create Users
func (repo *repository) CreateCreator(ctx context.Context, cons Creator) (Creator, error) {
	response, err := repo.queries.CreateCreator(ctx, CreateCreatorParams(cons))

	if err != nil {
		fmt.Print(err)
	}
	return response, err
}

func (repo *repository) Close() error {
	return repo.db.Close()
}

func NewUserStore(dbURL string) (Repository, error) {
	db, err := database.EstablishConnection(dbURL)
	if err != nil {
		fmt.Println("Error while establishing connection with databse " + err.Error())
	}

	return &repository{
		db:      db,
		queries: New(db),
	}, nil
}
