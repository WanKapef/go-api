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

	router.HandleFunc("/users", userHandler.Create).Methods("POST")
	router.HandleFunc("/users", userHandler.List).Methods("GET")
	router.HandleFunc("/users", userHandler.Update).Methods("PUT")
	router.HandleFunc("/users/{id}", httpx.WithID(userHandler.Delete)).Methods("DELETE")

	log.Println("API rodando na porta", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}
