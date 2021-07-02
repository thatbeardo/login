package clients_test

import (
	"net/http"
	"testing"

	"github.com/fanfit/login/testutil"
)

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
