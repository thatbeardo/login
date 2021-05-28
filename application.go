// @title Fan fit login
// @version 0.1.0

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @description ## Users
//@description ---
//
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
//
// @BasePath /
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fanfit/userservice/api/handlers/users"
	"github.com/fanfit/userservice/api/middleware"
	"github.com/fanfit/userservice/models/users/repository"
	"github.com/fanfit/userservice/models/users/service"
	"github.com/sirupsen/logrus"

	"github.com/fanfit/userservice/server"
	"github.com/gin-gonic/gin"
)

func main() {
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{
		TimestampFormat: time.StampMilli,
		FullTimestamp:   true,
	}

	db, err := server.CreatePostGresConnection(log, os.Getenv("DB_URL"))
	if err != nil {
		fmt.Print("Something went wrong!" + err.Error())
	}

	userRepository := repository.NewUserStore(db)
	userService := service.New(userRepository)

	engine := gin.Default()
	router := server.GenerateRouter(engine)

	router.Use(middleware.VerifyToken)
	users.Routes(router, userService)

	server.Orchestrate(engine, db)
}
