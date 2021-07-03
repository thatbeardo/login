package clients

import (
	"net/http"

	"github.com/fanfit/user-service/api/views"
	"github.com/fanfit/user-service/models/clients/service"
	"github.com/gin-gonic/gin"
)

// @Summary Get client and user fields by Email
// @Tags Clients
// @Description Get a client
// @Accept  json
// @Produce  json
// @Param email_id path string true "User Email ID"
// @Success 200 {object} repository.GetClientByEmailRow	"ok"
// @Success 404 {object} views.ErrView
// @Success 500 {object} views.ErrView
// @Security ApiKeyAuth
// @Router /v1/clients/{email_id} [get]
func getClientByEmail(service service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		emailID := c.Param("email_id")
		resource, err := service.GetClientByEmail(c.Request.Context(), emailID)
		if err != nil {
			views.Wrap(err, c)
			return
		}
		c.JSON(http.StatusOK, resource)
	}
}
