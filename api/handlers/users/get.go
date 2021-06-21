package users

import (
	"net/http"

	"github.com/fanfit/login/api/views"
	"github.com/fanfit/login/models/users/service"

	"github.com/gin-gonic/gin"
)

// @Summary Get user by Email ID
// @Tags Users
// @Description Get a user by its Email ID
// @Accept  json
// @Produce  json
// @Param email_id path string true "User Email ID"
// @Success 200 {object} repository.User	"ok"
// @Success 404 {object} views.ErrView
// @Success 500 {object} views.ErrView
// @Security ApiKeyAuth
// @Router /v1/users/{email_id} [get]
func getByID(service service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		emailID := c.Param("email_id")
		resource, err := service.GetByEmail(c.Request.Context(), emailID)
		if err != nil {
			views.Wrap(err, c)
			return
		}
		c.JSON(http.StatusOK, resource)
	}
}

func getCreatorByID(service service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		emailID := c.Param("email_id")
		resource, err := service.GetCreator(c.Request.Context(), emailID)
		if err != nil {
			views.Wrap(err, c)
			return
		}
		c.JSON(http.StatusOK, resource)
	}
}

func getClientByID(service service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		emailID := c.Param("email_id")
		resource, err := service.GetClient(c.Request.Context(), emailID)
		if err != nil {
			views.Wrap(err, c)
			return
		}
		c.JSON(http.StatusOK, resource)
	}
}
