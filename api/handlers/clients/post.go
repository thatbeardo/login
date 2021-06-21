package clients

import (
	"fmt"
	"net/http"

	"github.com/fanfit/login/api/views"
	"github.com/fanfit/login/models/clients/repository"
	"github.com/fanfit/login/models/clients/service"
	"github.com/gin-gonic/gin"
)

// @Summary Create a new Client
// @Description Adds data to 2 tables: Adds a user to the user table, and a client to the client table using the same user_id
// @Tags Clients
// @Accept  json
// @Produce  json
// @Param input body repository.User true "Details of the new user"
// @Success 202 {object} repository.GetClientByIDRow	"ok"
// @Failure 500 {object} views.ErrView	"ok"
// @Security ApiKeyAuth
// @Router /v1/users/ [post]
func post(service service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input repository.Client
		if err := c.ShouldBind(&input); err != nil {
			views.Wrap(err, c)
			return
		}

		fmt.Println("About to create")
		completeClient, err := service.CreateClient(c.Request.Context(), input)
		if err != nil {
			views.Wrap(err, c)
			return
		}

		c.JSON(http.StatusAccepted, completeClient)
	}
}
