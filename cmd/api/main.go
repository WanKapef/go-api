// @title           Go API
// @version         1.0
// @description     API REST com Gorilla Mux, SQLite, JWT e Swagger.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Suporte
// @contact.url    http://github.com/WanKapef
// @contact.email  suporte@wan.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @host      localhost:8080
// @BasePath  /

package main

import (
	"log"
	"net/http"

	"github.com/WanKapef/go-api/internal/config"
	"github.com/WanKapef/go-api/internal/database"
	"github.com/WanKapef/go-api/internal/handler"
	"github.com/WanKapef/go-api/internal/httpx"
	"github.com/WanKapef/go-api/internal/middleware"
	"github.com/WanKapef/go-api/internal/repository"
	"github.com/WanKapef/go-api/internal/service"

	_ "github.com/WanKapef/go-api/docs"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	cfg := config.Load()

	// conecta banco
	db := database.ConnectSQLite(cfg.DatabasePath)
	defer db.Close()

	// instancia dependências
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	authHandler := handler.NewAuthHandler(userService)

	router := mux.NewRouter()

	// Swagger
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)

	router.Handle("/login", middleware.ErrorMiddleware(authHandler.Login)).Methods("POST")
	router.Handle("/me", middleware.ErrorMiddleware(middleware.JWT(authHandler.Me))).Methods("GET")

	router.Handle("/users", middleware.ErrorMiddleware(userHandler.Create)).Methods("POST")
	router.Handle("/users", middleware.ErrorMiddleware(userHandler.List)).Methods("GET")
	router.Handle("/users/{id}", middleware.ErrorMiddleware(httpx.WithID(userHandler.GetByID))).Methods("GET")
	router.Handle("/users", middleware.ErrorMiddleware(userHandler.Update)).Methods("PUT")
	router.Handle("/users/{id}", middleware.ErrorMiddleware(httpx.WithID(userHandler.Delete))).Methods("DELETE")

	log.Println("API rodando na porta", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}
