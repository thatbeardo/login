package clients

import (
	"github.com/fanfit/userservice/models/clients/service"
	"github.com/gin-gonic/gin"
)

// Routes sets up resource specific routes on the engine instance
func Routes(r *gin.RouterGroup, service service.Service) {
	router := r.Group("/users")
	router.GET("/:email_id", getByID(service))
	router.POST("/", post(service))
}
