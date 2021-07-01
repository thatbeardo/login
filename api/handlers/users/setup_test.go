package users_test

import (
	"context"

	handler "github.com/fanfit/login/api/handlers"
	"github.com/fanfit/login/api/handlers/users"
	"github.com/fanfit/login/models/users/repository"
	"github.com/fanfit/login/models/users/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type mockService struct {
	GetByIDResponse repository.User
	CreateResponse  repository.GetClientByIDRow

	GetByIDErr error
	CreateErr  error
	DeleteErr  error
}

func (m mockService) GetByEmail(ctx context.Context, id string) (repository.User, error) {
	return m.GetByIDResponse, m.GetByIDErr
}

func (m mockService) Create(ctx context.Context, user repository.User) (repository.GetClientByIDRow, error) {
	return m.CreateResponse, m.CreateErr
}

func (m mockService) Delete(ctx context.Context, id string) error {
	return m.DeleteErr
}

func setupRouter(s service.Service) *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.NoRoute(handler.NoRoute)
	group := r.Group("/v1")
	users.Routes(group, s)
	return r
}
