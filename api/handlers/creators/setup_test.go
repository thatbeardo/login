package creators_test

import (
	"context"

	handler "github.com/fanfit/login/api/handlers"
	"github.com/fanfit/login/api/handlers/creators"
	"github.com/fanfit/login/models/creators/repository"
	"github.com/fanfit/login/models/creators/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

func setupRouter(s service.Service) *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.NoRoute(handler.NoRoute)
	group := r.Group("/v1")
	creators.Routes(group, s)
	return r
}
