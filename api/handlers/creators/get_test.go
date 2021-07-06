package creators_test

import (
	"database/sql"
	"net/http"
	"testing"

	handler "github.com/fanfit/user-service/api/handlers"
	"github.com/fanfit/user-service/api/handlers/creators"
	"github.com/fanfit/user-service/models/creators/repository"
	"github.com/fanfit/user-service/models/creators/service"
	"github.com/fanfit/user-service/testutil"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var testUser = repository.GetCreatorByEmailRow{
	ID:                0,
	UserType:          "creator",
	FirstName:         "Scott",
	LastName:          "Mathison",
	Email:             "gdalessa@usc.edu",
	Username:          sql.NullString{},
	PhoneNo:           sql.NullString{},
	Gender:            sql.NullString{},
	ProfilePicture:    sql.NullString{},
	Bio:               sql.NullString{},
	FanfitUserID:      0,
	PaymentInfo:       sql.NullString{},
	LogoPicture:       sql.NullString{},
	BackgroundPicture: sql.NullString{},
}

func TestHandler_InvalidPathCalled_StatusNotFound(t *testing.T) {
	mockService := mockService{}

	router := setupRouter(mockService)
	response, cleanup := testutil.PerformRequest(router, "GET", "/v1/invalid-path/", "")
	defer cleanup()

	testutil.ValidateResponse(t, response, testutil.GenerateError("", "query-parameter-todo", "Path not found", http.StatusNotFound), http.StatusNotFound)
}
func TestHandler_ValidPathCalled_StatusOK(t *testing.T) {
	mockService := mockService{GetCreatorByEmailResponse: testUser}

	router := setupRouter(mockService)
	response, cleanup := testutil.PerformRequest(router, "GET", "/v1/creators/gdalessa@usc.edu", "")
	defer cleanup()

	testutil.ValidateResponse(t, response, testUser, http.StatusOK)
}

func setupRouter(s service.Service) *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.NoRoute(handler.NoRoute)
	group := r.Group("/v1")
	creators.Routes(group, s)
	return r
}
