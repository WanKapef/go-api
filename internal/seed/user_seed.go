package seed

import (
	"database/sql"
	"log"

	"github.com/go-faker/faker/v4"
)

func UserSeed(db *sql.DB) error {
	for i := 0; i < 10; i++ {
		name := faker.Name()
		email := faker.Email()

		var exists int
		err := db.QueryRow(
			"SELECT COUNT(1) FROM users WHERE email = ?",
			email,
		).Scan(&exists)

		if err != nil {
			return err
		}

		if exists > 0 {
			continue // já existe idempotência
		}

		_, err = db.Exec(
			"INSERT INTO users (name, email) VALUES (?, ?)",
			name, email,
		)
		if err != nil {
			return err
		}
	}

	log.Println("👤 Usuários seedados")
	return nil
}
