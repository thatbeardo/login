package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

// Repository is used by the service to communicate with the underlying database
type Repository interface {
	// CREATES

	// DELETES
	Delete(context.Context, string) error

	// GETS
	GetByEmail(context.Context, string) (User, error)

	Create(ctx context.Context, user User) (GetClientByIDRow, error)
}

type repository struct {
	queries *Queries
	db      *sql.DB
}

// GetByID function adds a resource node
func (repo *repository) GetByEmail(ctx context.Context, email string) (User, error) {
	transaction, err := repo.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		fmt.Print("somethinbg went wrong")
	}

	response, err := repo.queries.WithTx(transaction).GetUserByEmail(ctx, email)
	if err != nil {
		transaction.Rollback()
		fmt.Print(err)
	}

	response, err = repo.queries.WithTx(transaction).GetUserByEmail(ctx, email)
	if err != nil {
		transaction.Rollback()
		fmt.Print(err)
	}

	// Commit the change if all queries ran successfully
	err = transaction.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return response, err
}

// Create function adds a node to the graph - typically invoked by customer API not guard-my-app
func (repo *repository) Create(ctx context.Context, user User) (GetClientByIDRow, error) {
	transaction, err := repo.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		fmt.Print("somethinbg went wrong")
	}

	response, err := repo.queries.WithTx(transaction).CreateUser(ctx, CreateUserParams{
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Email:          user.Email,
		UserTypeID:     user.UserTypeID,
		Username:       user.Username,
		PhoneNo:        user.PhoneNo,
		Gender:         user.Gender,
		ProfilePicture: user.ProfilePicture,
		Bio:            user.Bio,
	})
	if err != nil {
		transaction.Rollback()
		fmt.Print(err)
	}
	_, err = repo.queries.WithTx(transaction).CreateClient(ctx, CreateClientParams{
		FanfitUserID: response.ID,
		TempField:    sql.NullString{},
	})
	if err != nil {
		transaction.Rollback()
		fmt.Print(err)
	}

	// Commit the change if all queries ran successfully
	err = transaction.Commit()
	if err != nil {
		log.Fatal(err)
	}
	fullClientobj, err := repo.queries.WithTx(transaction).GetClientByID(ctx, response.ID)
	return fullClientobj, err
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
