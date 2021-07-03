// @title Fan fit user-service
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
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"

	// API Routes

	creatorHandlers "github.com/fanfit/user-service/api/handlers/creators"
	userHandlers "github.com/fanfit/user-service/api/handlers/users"

	// Tags
	// Users Tag
	userRepository "github.com/fanfit/user-service/models/users/repository"
	userServicePackage "github.com/fanfit/user-service/models/users/service"

	creatorRepository "github.com/fanfit/user-service/models/creators/repository"
	creatorServicePackage "github.com/fanfit/user-service/models/creators/service"

	// Creators Tag

	// Clients Tag

	"github.com/fanfit/user-service/server"
	"github.com/gin-gonic/gin"
)

type envVars struct {
	dbUserName string
	dbPassword string
	dbHost     string
	dbPort     string
	dbName     string
	dbSchema   string
}

func main() {
	envVars, err := loadEnvVars()
	if err != nil {
		log.Fatal(fmt.Sprintf("Error while loading Env variables: %s", err.Error()))
	}
	dbURL := prepareDbURL(envVars)

	// Instantiate service for each tag
	userStore, err := userRepository.NewUserStore(dbURL)
	if err != nil {
		log.Fatal(fmt.Printf("Error while creating userStore: %s", err.Error()))
	}
	userService := userServicePackage.New(userStore)

	creatorStore, err := creatorRepository.NewCreatorStore(dbURL)
	if err != nil {
		fmt.Printf("Error while creating userStore: %s", err.Error())
		os.Exit(1)
	}
	creatorService := creatorServicePackage.New(creatorStore)

	// clientStore := clientRepository.NewUserStore(db)
	// clientService := clientServicePackage.New(clientStore)

	// Initialize the middleware and routes
	engine := gin.Default()
	router := server.GenerateRouter(engine)

	// Set routes for each tag
	// router.Use(middleware.VerifyToken)
	userHandlers.Routes(router, userService)
	creatorHandlers.Routes(router, creatorService)
	// clientHandlers.Routes(router, clientService)

	server.Orchestrate(engine, userStore, creatorStore)
}

func loadEnvVars() (*envVars, error) {
	dbUsername, envPresent := os.LookupEnv("DB_USERNAME")
	if !envPresent {
		return nil, errors.New("DB_USERNAME environment variable missing")
	}

	dbPassword, envPresent := os.LookupEnv("DB_PASSWORD")
	if !envPresent {
		return nil, errors.New("DB_PASSWORD environment variable missing")
	}

	dbHost, envPresent := os.LookupEnv("DB_HOST")
	if !envPresent {
		return nil, errors.New("DB_HOST environment variable missing")
	}

	dbPort, envPresent := os.LookupEnv("DB_PORT")
	if !envPresent {
		return nil, errors.New("DB_PORT environment variable missing")
	}

	dbName, envPresent := os.LookupEnv("DB_NAME")
	if !envPresent {
		return nil, errors.New("DB_NAME environment variable missing")
	}

	dbSchema, envPresent := os.LookupEnv("DB_SCHEMA")
	if !envPresent {
		return nil, errors.New("DB_SCHEMA environment variable missing")
	}

	return &envVars{
		dbUserName: dbUsername,
		dbPassword: dbPassword,
		dbHost:     dbHost,
		dbPort:     dbPort,
		dbName:     dbName,
		dbSchema:   dbSchema,
	}, nil
}

func prepareDbURL(envVars *envVars) string {
	dbURL := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(envVars.dbUserName, envVars.dbPassword),
		Host:   fmt.Sprintf("%s:%s", envVars.dbHost, envVars.dbPort),
		Path:   envVars.dbName,
	}

	q := dbURL.Query()
	q.Add("search_path", envVars.dbSchema)
	dbURL.RawQuery = q.Encode()
	return dbURL.String()
}
