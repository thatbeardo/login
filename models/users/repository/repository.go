package repository

import (
	"context"
	"fmt"
)

// Repository is used by the service to communicate with the underlying database
type Repository interface {
	// CREATES
	Create(context.Context, User) (User, error)

	// DELETES
	Delete(context.Context, string) error

	// GETS
	GetByEmail(context.Context, string) (User, error)
}

type repository struct {
	queries *Queries
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
