package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/fanfit/user-service/database"
)

// Repository is used by the service to communicate with the underlying database
type Repository interface {
	// CREATES

	// DELETES
	Delete(context.Context, string) error

	// GETS
	GetByEmail(context.Context, string) (User, error)

	Create(ctx context.Context, user User) (GetClientByIDRow, error)

	Close() error
}

type repository struct {
	queries *Queries
	db      *sql.DB
}

// GetByID function adds a resource node
func (repo *repository) GetByEmail(ctx context.Context, email string) (User, error) {
	response, err := repo.queries.GetUserByEmail(ctx, email)
	if err != nil {
		fmt.Print(err)
	}

	return response, err
}

// Create function adds a node to the graph - typically invoked by customer API not guard-my-app
func (repo *repository) Create(ctx context.Context, user User) (GetClientByIDRow, error) {
	transaction, err := repo.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		fmt.Print("something went wrong while executing transaction: " + err.Error())
	}

	response, err := repo.queries.WithTx(transaction).CreateUser(ctx, CreateUserParams{
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Email:          user.Email,
		UserType:       user.UserType,
		Username:       user.Username,
		PhoneNo:        user.PhoneNo,
		Gender:         user.Gender,
		ProfilePicture: user.ProfilePicture,
		Bio:            user.Bio,
	})
	if err != nil {
		transaction.Rollback()
		return GetClientByIDRow{}, err
	}
	_, err = repo.queries.WithTx(transaction).CreateClient(ctx, CreateClientParams{
		FanfitUserID: response.ID,
		TempField:    sql.NullString{},
	})
	if err != nil {
		transaction.Rollback()
		return GetClientByIDRow{}, err
	}

	fullClientobj, err := repo.queries.WithTx(transaction).GetClientByID(ctx, response.ID)
	if err != nil {
		transaction.Rollback()
		return GetClientByIDRow{}, err
	}

	// Commit the change if all queries ran successfully
	err = transaction.Commit()
	if err != nil {
		return GetClientByIDRow{}, err
	}
	return fullClientobj, err
}

// Delete function deletes a node from the graph
func (repo *repository) Delete(ctx context.Context, id string) error {
	return nil
}

// Delete function deletes a node from the graph
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
