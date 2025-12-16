package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
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

	return &Config{
		Port: port,
	}
}
