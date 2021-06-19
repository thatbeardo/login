package creators_test

import (
	"context"

	"github.com/fanfit/login/models/creators/repository"
)

type mockService struct {
	GetCreatorResponse    repository.GetCreatorRow
	CreateCreatorResponse repository.Creator

	GetClientError  error
	CreateClientErr error
}

func (m mockService) GetCreator(ctx context.Context, id string) (repository.GetCreatorRow, error) {
	return m.GetCreatorResponse, m.GetClientError
}

func (m mockService) CreateCreator(ctx context.Context, user repository.Creator) (repository.Creator, error) {
	return m.CreateCreatorResponse, m.CreateClientErr
}
