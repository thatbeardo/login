package clients_test

import (
	"database/sql"
	"net/http"
	"testing"

	"github.com/fanfit/user-service/models/clients/repository"
	"github.com/fanfit/user-service/testutil"
)

var payload = "{\"ID\":0,\"UserType\":\"creator\",\"FirstName\":\"Gabe\",\"LastName\":\"\",\"Email\":\"\",\"CreatedDate\":\"0001-01-01T00:00:00Z\",\"Username\":{\"String\":\"\",\"Valid\":false},\"PhoneNo\":{\"String\":\"\",\"Valid\":false},\"Gender\":{\"String\":\"\",\"Valid\":false},\"ProfilePicture\":{\"String\":\"\",\"Valid\":false},\"Bio\":{\"String\":\"\",\"Valid\":false},\"FanfitUserID\":0,\"TempField\":{\"String\":\"\",\"Valid\":false}}"

var testUser2 = repository.GetClientByIDRow{
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

func TestHandler_InvalidPayload_StatusNotFound(t *testing.T) {
	mockService := mockService{}
	router := setupRouter(mockService)
	response, cleanup := testutil.PerformRequest(router, "POST", "/v1/invalid-path/", "")
	defer cleanup()
	testutil.ValidateResponse(t, response, testutil.GenerateError("", "query-parameter-todo", "Path not found", http.StatusNotFound), http.StatusNotFound)

}

func TestHandler_InvalidPayload_StatusInternalServerError(t *testing.T) {
	mockService := mockService{}
	router := setupRouter(mockService)
	response, cleanup := testutil.PerformRequest(router, "POST", "/v1/clients/", "")
	defer cleanup()
	testutil.ValidateResponse(t, response, testutil.GenerateError("/v1/clients/", "query-parameter-todo", "EOF", http.StatusInternalServerError), http.StatusInternalServerError)

}

func TestHandler_ValidPayload_StatusAccepted(t *testing.T) {
	mockService := mockService{CreateClientResponse: testUser2}
	router := setupRouter(mockService)
	response, cleanup := testutil.PerformRequest(router, "POST", "/v1/clients/", payload)
	defer cleanup()
	testutil.ValidateResponse(t, response, testUser2, http.StatusAccepted)

}
