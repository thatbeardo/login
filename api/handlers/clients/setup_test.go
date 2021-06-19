package clients_test

import (
	"context"

	"github.com/fanfit/login/models/clients/repository"
)

type mockService struct {
	GetClientResponse    repository.GetClientByEmailRow
	CreateClientResponse repository.Client

	GetClientError  error
	CreateClientErr error
}

func (m mockService) GetClientByEmail(ctx context.Context, id string) (repository.GetClientByEmailRow, error) {
	return m.GetClientResponse, m.GetClientError
}

func (m mockService) CreateClient(ctx context.Context, user repository.Client) (repository.Client, error) {
	return m.CreateClientResponse, m.CreateClientErr
}
