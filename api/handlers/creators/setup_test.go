package creators_test

import (
	"context"

	"github.com/fanfit/login/models/creators/repository"
)

type mockService struct {
	GetCreatorByEmailResponse repository.GetCreatorByEmailRow
	CreateCreatorResponse     repository.Creator

	GetClientError  error
	CreateClientErr error
}

func (m mockService) GetCreatorByEmail(ctx context.Context, id string) (repository.GetCreatorByEmailRow, error) {
	return m.GetCreatorByEmailResponse, m.GetClientError
}

func (m mockService) CreateCreator(ctx context.Context, user repository.Creator) (repository.Creator, error) {
	return m.CreateCreatorResponse, m.CreateClientErr
}
