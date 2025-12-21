package main

import (
	"log"

	"github.com/WanKapef/go-api/internal/config"
	"github.com/WanKapef/go-api/internal/database"
	"github.com/WanKapef/go-api/internal/seed"
)

func main() {
	cfg := config.Load()

	db := database.ConnectSQLite(cfg.DatabasePath)
	defer db.Close()

	log.Println("🌱 Rodando seeds...")

	if err := seed.Run(db); err != nil {
		log.Fatal(err)
	}

	log.Println("✅ Seeds finalizados")
}
