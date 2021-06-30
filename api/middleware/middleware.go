package middleware

import (
	"net/http"

	"github.com/fanfit/user-service/api/middleware/injection"
	"github.com/fanfit/user-service/api/views"
	"github.com/gin-gonic/gin"
)

// VerifyToken function verifies the incoming jwt token
func VerifyToken(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, x--tenant")
	if err := injection.VerifyAccessToken(c.Writer, c.Request); err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			views.GenerateErrorResponse(
				http.StatusUnauthorized,
				"The access token is invalid. Please provide a valid token in the header",
				c.Request.URL.Path,
			),
		)
	}
}
