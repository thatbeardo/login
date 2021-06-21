package service

import (
	"context"
	"fmt"

	"github.com/fanfit/login/models/clients/repository"
)

// Service receives commands from handlers and forwards them to the repository
type Service interface {
	GetClientByEmail(context.Context, string) (repository.GetClientByEmailRow, error)
	CreateClient(context.Context, repository.Client) (repository.GetClientByIDRow, error)
}

type service struct {
	repository repository.Repository
}

// New creates a service instance with the repository passed
func New(repository repository.Repository) Service {
	return &service{repository: repository}
}

func (service *service) GetClientByEmail(ctx context.Context, input string) (repository.GetClientByEmailRow, error) {
	return service.repository.GetClientByEmail(ctx, input)
}

func (service *service) CreateClient(ctx context.Context, input repository.Client) (repository.GetClientByIDRow, error) {
	fmt.Print("Going into repo")
	return service.repository.CreateClient(ctx, input)
}
