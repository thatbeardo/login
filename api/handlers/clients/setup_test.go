package clients_test

import (
	"context"

	"github.com/fanfit/login/models/clients/repository"
)

type mockService struct {
	GetClientResponse    repository.GetClientsRow
	CreateClientResponse repository.Client

	GetClientError  error
	CreateClientErr error
}

func (m mockService) GetClients(ctx context.Context, id string) (repository.GetClientsRow, error) {
	return m.GetClientResponse, m.GetClientError
}

func (m mockService) CreateClients(ctx context.Context, user repository.Client) (repository.Client, error) {
	return m.CreateClientResponse, m.CreateClientErr
}
