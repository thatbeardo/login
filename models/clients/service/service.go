package service

import (
	"context"
	"fmt"

	"github.com/fanfit/login/models/clients/repository"
)

// Service receives commands from handlers and forwards them to the repository
type Service interface {
	GetClients(context.Context, string) (repository.GetClientsRow, error)
	CreateClients(context.Context, repository.Client) (repository.Client, error)
}

type service struct {
	repository repository.Repository
}

// New creates a service instance with the repository passed
func New(repository repository.Repository) Service {
	return &service{repository: repository}
}

func (service *service) GetClients(ctx context.Context, input string) (repository.GetClientsRow, error) {
	return service.repository.GetClients(ctx, input)
}

func (service *service) CreateClients(ctx context.Context, input repository.Client) (repository.Client, error) {
	fmt.Print("Going into repo")
	return service.repository.CreateClients(ctx, input)
}
