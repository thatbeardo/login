package clients_test

import (
	"database/sql"
	"net/http"
	"testing"

	"github.com/fanfit/login/models/clients/repository"
	"github.com/fanfit/login/testutil"
)

var testUser = repository.GetClientByEmailRow{
	ID:             0,
	UserTypeID:     0,
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
	response, cleanup := testutil.PerformRequest(router, "GET", "/v1/clients/gdalessausc.edu", "")
	defer cleanup()
	testutil.ValidateResponse(t, response, testUser, http.StatusOK)

}
