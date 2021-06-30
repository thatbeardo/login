package creators

import (
	"github.com/fanfit/login/models/creators/service"
	"github.com/gin-gonic/gin"
)

// Routes sets up resource specific routes on the engine instance
func Routes(r *gin.RouterGroup, service service.Service) {
	router := r.Group("/creators")
	router.GET("/:email_id", getCreatorByEmail(service))
	router.POST("/", post(service))
}
