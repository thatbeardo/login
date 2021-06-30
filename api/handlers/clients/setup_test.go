package clients_test

import (
	"context"

	"github.com/fanfit/user-service/models/clients/repository"
)

type mockService struct {
	GetClientResponse    repository.GetClientByEmailRow
	CreateClientResponse repository.GetClientByIDRow

	GetClientError  error
	CreateClientErr error
}

func (m mockService) GetClientByEmail(ctx context.Context, id string) (repository.GetClientByEmailRow, error) {
	return m.GetClientResponse, m.GetClientError
}

func (m mockService) CreateClient(ctx context.Context, client repository.Client) (repository.GetClientByIDRow, error) {
	return m.CreateClientResponse, m.CreateClientErr
}
