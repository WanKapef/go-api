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

	"github.com/gorilla/mux"
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

	router := mux.NewRouter()

	router.Use(middleware.Logger)

	router.Handle("/users", middleware.ErrorMiddleware(userHandler.Create)).Methods("POST")
	router.Handle("/users", middleware.ErrorMiddleware(userHandler.List)).Methods("GET")
	router.Handle("/users/{id}", middleware.ErrorMiddleware(httpx.WithID(userHandler.ListByID))).Methods("GET")
	router.Handle("/users", middleware.ErrorMiddleware(userHandler.Update)).Methods("PUT")
	router.Handle("/users/{id}", middleware.ErrorMiddleware(httpx.WithID(userHandler.Delete))).Methods("DELETE")

	log.Println("API rodando na porta", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}
