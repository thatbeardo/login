package clients_test

import (
	"database/sql"
	"net/http"
	"testing"

	handler "github.com/fanfit/user-service/api/handlers"
	"github.com/fanfit/user-service/api/handlers/clients"
	"github.com/fanfit/user-service/models/clients/repository"
	"github.com/fanfit/user-service/models/clients/service"
	"github.com/fanfit/user-service/testutil"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var testUser = repository.GetClientByEmailRow{
	ID:             0,
	UserType:       "creator",
	FirstName:      "",
	LastName:       "",
	Email:          "",
	Username:       sql.NullString{},
	PhoneNo:        sql.NullString{},
	Gender:         sql.NullString{},
	ProfilePicture: sql.NullString{},
	Bio:            sql.NullString{},
	FanfitUserID:   0,
	TempField:      sql.NullString{},
}

func TestHandler_InvalidPathCalled_StatusNotFound(t *testing.T) {
	mockService := mockService{}

	router := setupRouter(mockService)
	response, cleanup := testutil.PerformRequest(router, "GET", "/v1/invalid-path/", "")
	defer cleanup()

	testutil.ValidateResponse(t, response, testutil.GenerateError("", "query-parameter-todo", "Path not found", http.StatusNotFound), http.StatusNotFound)
}

func TestHandler_ValidPathCalled_StatusOK(t *testing.T) {
	mockService := mockService{GetClientResponse: testUser}
	router := setupRouter(mockService)
	response, cleanup := testutil.PerformRequest(router, "GET", "/v1/clients/gdalessa@usc.edu", "")
	defer cleanup()
	testutil.ValidateResponse(t, response, testUser2, http.StatusOK)

}

func setupRouter(s service.Service) *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.NoRoute(handler.NoRoute)
	group := r.Group("/v1")
	clients.Routes(group, s)
	return r
}
