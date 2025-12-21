package seed

import "database/sql"

func Run(db *sql.DB) error {
	if err := UserSeed(db); err != nil {
		return err
	}
	return nil
}
