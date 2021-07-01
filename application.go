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
	"errors"
	"fmt"
	"net/url"
	"os"

	// API Routes

	userHandlers "github.com/fanfit/login/api/handlers/users"

	// Tags
	// Users Tag
	userRepository "github.com/fanfit/login/models/users/repository"
	userServicePackage "github.com/fanfit/login/models/users/service"

	// Creators Tag

	// Clients Tag

	"github.com/fanfit/login/api/middleware"
	"github.com/fanfit/login/server"
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
		fmt.Printf("Error while loading Env variables: %s", err.Error())
	}

	dbURL := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(envVars.dbUserName, envVars.dbPassword),
		Host:   fmt.Sprintf("%s:%s", envVars.dbHost, envVars.dbPort),
		Path:   envVars.dbName,
	}

	q := dbURL.Query()
	q.Add("schema", envVars.dbSchema)
	encodedURL := q.Encode()

	// dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("DB_USERNAME"), , , os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	// db, err := server.CreatePostGresConnection(log, dbURL)
	// if err != nil {
	// 	fmt.Print("Something went wrong!" + err.Error())
	// }

	// Instantiate service for each tag
	userStore, err := userRepository.NewUserStore(encodedURL)
	if err != nil {
		fmt.Printf("Error while creating userStore: %s", err.Error())
	}
	userService := userServicePackage.New(userStore)

	// creatorStore := creatorRepository.NewUserStore(db)
	// creatorService := creatorServicePackage.New(creatorStore)

	// clientStore := clientRepository.NewUserStore(db)
	// clientService := clientServicePackage.New(clientStore)

	// Initialize the middleware and routes
	engine := gin.Default()
	router := server.GenerateRouter(engine)

	// Set routes for each tag
	router.Use(middleware.VerifyToken)
	userHandlers.Routes(router, userService)
	// clientHandlers.Routes(router, clientService)
	// creatorHandlers.Routes(router, creatorService)

	server.Orchestrate(engine, userStore)
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
