package repository

import (
	"context"
	"database/sql"
	"fmt"
)

// Repository is used by the service to communicate with the underlying database
type Repository interface {
	// CREATES
	Create(context.Context, User) (User, error)
	CreateConsumer(context.Context, int32) (Consumer, error)

	// DELETES
	Delete(context.Context, string) error

	// GETS
	GetByEmail(context.Context, string) (User, error)
	GetCreator(context.Context, string) (GetCreatorRow, error)
	GetClient(context.Context, string) (GetClientRow, error)
}

type repository struct {
	queries *Queries
}

// GetCreator with fan_fit_userid
func (repo *repository) GetCreator(ctx context.Context, emailID string) (GetCreatorRow, error) {
	response, err := repo.queries.GetCreator(ctx, emailID)

	if err != nil {
		fmt.Print(err)
	}

	return response, err

}

// GetConsumer
func (repo *repository) GetClient(ctx context.Context, emailID string) (GetClientRow, error) {
	response, err := repo.queries.GetClient(ctx, emailID)
	if err != nil {
		fmt.Print(err)
	}
	return response, err
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
func (repo *repository) Create(ctx context.Context, user User) (User, error) {
	response, err := repo.queries.CreateUser(ctx, CreateUserParams{
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
func (repo *repository) CreateConsumer(ctx context.Context, userID int32) (Consumer, error) {
	var temp sql.NullString
	temp.String = "blank"
	temp.Valid = true

	response, err := repo.queries.CreateConsumer(ctx, CreateConsumerParams{
		FanfitUserID: userID,
		TempField:    temp,
	})

	if err != nil {
		fmt.Print(err)
	}
	return response, err
}
