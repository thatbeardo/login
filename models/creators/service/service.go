package service

import (
	"context"
	"fmt"

	"github.com/fanfit/userservice/models/creators/repository"
)

// Service receives commands from handlers and forwards them to the repository
type Service interface {
	GetCreator(context.Context, int32) (repository.GetCreatorRow, error)
	CreateCreator(context.Context, repository.Creator) (repository.Creator, error)
}

type service struct {
	repository repository.Repository
}

// New creates a service instance with the repository passed
func New(repository repository.Repository) Service {
	return &service{repository: repository}
}

func (service *service) GetCreator(ctx context.Context, input int32) (repository.GetCreatorRow, error) {
	return service.repository.GetCreator(ctx, input)
}

func (service *service) CreateCreator(ctx context.Context, input repository.Creator) (repository.Creator, error) {
	fmt.Print("Going into repo")
	return service.repository.CreateCreator(ctx, input)
}
