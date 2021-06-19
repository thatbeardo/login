package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/jackc/pgx/stdlib"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/fanfit/login/docs"
)

// GenerateRouter instantiates and initializes a new Router.
func GenerateRouter(r *gin.Engine) *gin.RouterGroup {
	config := cors.DefaultConfig()
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")
	config.AllowMethods = append(config.AllowMethods, "OPTIONS")
	config.AllowAllOrigins = true
	r.Use(cors.New(config))
	setupSwagger(r)
	return r.Group("/v1")
}

func CreatePostGresConnection(logger *logrus.Logger, dbURL string) (*sql.DB, error) {

	// conn, err := pgx.Connect(context.Background(), dbURL)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	// 	os.Exit(1)
	// }

	// u, err := url.Parse(dbURL)
	// if err != nil {
	// 	return nil, err
	// }

	// c, err := pgx.ParseConfig(u.String())
	// if err != nil {
	// 	return nil, err
	// }

	db, err := sql.Open("pgx", os.Getenv("DB_URL"))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func setupSwagger(r *gin.Engine) {
	hostURL := fmt.Sprintf("http://%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	docs.SwaggerInfo.Host = hostURL
	r.StaticFile("/docs/swagger.json", "./docs/swagger.json")

	url := ginSwagger.URL(fmt.Sprintf("%s/docs/swagger.json", hostURL))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

// Orchestrate begins listening on PORT and gracefully shuts down the server incase of interrupt
func Orchestrate(router *gin.Engine, db *sql.DB) {
	srv := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shuting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	defer db.Close()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
