package middleware_test

import (
	"errors"
	"net/http"
	"testing"

	handler "github.com/fanfit/user-service/api/handlers"
	"github.com/fanfit/user-service/api/middleware"
	"github.com/fanfit/user-service/api/middleware/injection"
	"github.com/fanfit/user-service/testutil"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func TestVerifyJWT_VerificationFails_ReturnStatusUnauthorized(t *testing.T) {
	defer injection.Reset()
	injection.VerifyAccessToken = func(w gin.ResponseWriter, r *http.Request) error {
		return errors.New("test-error")
	}
	const path = "/test"
	router := setupRouter()
	router.GET(path, middleware.VerifyToken)
	response, cleanup := testutil.PerformRequest(router, "GET", path, "")
	defer cleanup()

	testutil.ValidateResponse(t, response,
		testutil.GenerateError(
			"/test", "query-parameter-todo",
			"The access token is invalid. Please provide a valid token in the header",
			401),
		http.StatusUnauthorized)
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.NoRoute(handler.NoRoute)
	return r
}
