package users

import (
	"fmt"
	"net/http"

	"github.com/fanfit/userservice/api/views"
	"github.com/fanfit/userservice/models/users/repository"
	"github.com/fanfit/userservice/models/users/service"
	"github.com/gin-gonic/gin"
)

// @Summary Create a new User
// @Description Add a new user to users table
// @Tags Users
// @Accept  json
// @Produce  json
// @Param input body repository.User true "Details of the new user"
// @Success 202 {object} repository.User	"ok"
// @Failure 500 {object} views.ErrView	"ok"
// @Security ApiKeyAuth
// @Router /v1/users/ [post]
func post(service service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input repository.User
		if err := c.ShouldBind(&input); err != nil {
			views.Wrap(err, c)
			return
		}
		fmt.Println("About to create")
		response, err := service.Create(c.Request.Context(), input)
		// 0 is for creator, 1 is for consumer
		userID, er := service.CreateConsumer(c.Request.Context(), response.ID)
		if userID != 0 && er != nil {
			views.Wrap(err, c)
			return
		}

		if err != nil {
			views.Wrap(err, c)
			return
		}
		c.JSON(http.StatusAccepted, response)
	}
}
