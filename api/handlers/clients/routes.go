package clients

import (
	"github.com/fanfit/user-service/models/clients/service"
	"github.com/gin-gonic/gin"
)

// Routes sets up resource specific routes on the engine instance
func Routes(r *gin.RouterGroup, service service.Service) {
	router := r.Group("/clients")
	router.GET("/:email_id", getClientByEmail(service))
	router.POST("/", post(service))
}
