package creators_test

import (
	"database/sql"
	"net/http"
	"testing"

	"github.com/fanfit/login/models/creators/repository"
	"github.com/fanfit/login/testutil"
)

var testUser = repository.GetCreatorByEmailRow{
	ID:                0,
	UserTypeID:        1,
	FirstName:         "",
	LastName:          "",
	Email:             "",
	Username:          sql.NullString{},
	PhoneNo:           sql.NullString{},
	Gender:            sql.NullString{},
	ProfilePicture:    sql.NullString{},
	Bio:               sql.NullString{},
	FanfitUserID:      0,
	PaymentInfo:       "",
	LogoPicture:       "",
	BackgroundPicture: "",
}

// Only doing get testing because
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
	response, cleanup := testutil.PerformRequest(router, "GET", "/v1/creators/gdalessausc.edu", "")
	defer cleanup()

	testutil.ValidateResponse(t, response, testUser, http.StatusOK)
}
