package service

import (
	"context"
	"fmt"

	"github.com/fanfit/userservice/models/users/repository"
)

// Service receives commands from handlers and forwards them to the repository
type Service interface {
	// CREATES
	Create(context.Context, repository.User) (repository.User, error)
	CreateConsumer(context.Context, int32) (repository.Consumer, error)

	// DELETES
	Delete(context.Context, string) error

	// GETS
	GetByEmail(context.Context, string) (repository.User, error)
	GetClient(context.Context, string) (repository.GetClientRow, error)
	GetCreator(context.Context, string) (repository.GetCreatorRow, error)
}

type service struct {
	repository repository.Repository
}

// New creates a service instance with the repository passed
func New(repository repository.Repository) Service {
	return &service{repository: repository}
}

func (service *service) GetClient(ctx context.Context, emailID string) (repository.GetClientRow, error) {
	return service.repository.GetClient(ctx, emailID)
}

func (service *service) GetCreator(ctx context.Context, email string) (repository.GetCreatorRow, error) {
	return service.repository.GetCreator(ctx, email)
}

func (service *service) GetByEmail(ctx context.Context, id string) (repository.User, error) {
	return service.repository.GetByEmail(ctx, id)
}

func (service *service) CreateConsumer(ctx context.Context, id int32) (repository.Consumer, error) {
	fmt.Print("Going into repo")
	return service.repository.CreateConsumer(ctx, id)
}

func (service *service) Create(ctx context.Context, input repository.User) (repository.User, error) {
	fmt.Print("Going into repo")
	return service.repository.Create(ctx, input)
}

func (service *service) Delete(ctx context.Context, id string) error {
	return service.repository.Delete(ctx, id)
}
