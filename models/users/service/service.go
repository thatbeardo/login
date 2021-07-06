package service

import (
	"context"

	"github.com/fanfit/user-service/models/users/repository"
)

// Service receives commands from handlers and forwards them to the repository
type Service interface {
	GetByEmail(context.Context, string) (repository.User, error)
	Create(context.Context, repository.User) (repository.GetClientByIDRow, error)
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

func (service *service) Create(ctx context.Context, input repository.User) (repository.GetClientByIDRow, error) {
	return service.repository.Create(ctx, input)
}

func (service *service) Delete(ctx context.Context, id string) error {
	return service.repository.Delete(ctx, id)
}
