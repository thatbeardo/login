package users

import (
	"github.com/fanfit/login/models/users/service"
	"github.com/gin-gonic/gin"
)

// Routes sets up resource specific routes on the engine instance
func Routes(r *gin.RouterGroup, service service.Service) {
	router := r.Group("/users")
	router.GET("/:email_id", getByID(service))
	router.GET("/creator/:email_id", getCreatorByID(service))
	router.GET("/client/:email_id", getClientByID(service))
	router.POST("/", post(service))
}
