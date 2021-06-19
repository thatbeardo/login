package service

import (
	"context"
	"fmt"

	"github.com/fanfit/login/models/users/repository"
)

// Service receives commands from handlers and forwards them to the repository
type Service interface {
	GetByEmail(context.Context, string) (repository.User, error)
	Create(context.Context, repository.User) (repository.User, error)
	Delete(context.Context, string) error
}

type service struct {
	repository repository.Repository
}

// New creates a service instance with the repository passed
func New(repository repository.Repository) Service {
	return &service{repository: repository}
}

func (service *service) GetByEmail(ctx context.Context, id string) (repository.User, error) {
	return service.repository.GetByEmail(ctx, id)
}

func (service *service) Create(ctx context.Context, input repository.User) (repository.User, error) {
	fmt.Print("Going into repo")
	return service.repository.Create(ctx, input)
}

func (service *service) Delete(ctx context.Context, id string) error {
	return service.repository.Delete(ctx, id)
}
