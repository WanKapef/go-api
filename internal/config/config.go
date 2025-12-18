package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	DatabasePath string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Nenhum arquivo .env encontrado, usando variáveis de ambiente do sistema")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Porta padrão
	}

	databasePath := os.Getenv("DATABASE_PATH")
	if databasePath == "" {
		log.Fatal("DATABASE_PATH não definido nas variáveis de ambiente")
	}

	return &Config{
		Port:         port,
		DatabasePath: databasePath,
	}
}
