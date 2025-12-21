package repository

import (
	"database/sql"
	"errors"

	"github.com/WanKapef/go-api/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create
func (r *UserRepository) Create(user *model.User) error {
	result, err := r.db.Exec(
		`INSERT INTO users (name, email) VALUES (?, ?)`,
		user.Name,
		user.Email,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = id
	return nil
}

// Read
func (r *UserRepository) List(limit, offset int, name, email, search string) ([]model.User, error) {

	query := `SELECT id, name, email FROM users WHERE 1=1`

	var args []any

	if name != "" {
		query += " AND name LIKE ?"
		args = append(args, "%"+name+"%")
	}

	if email != "" {
		query += " AND email LIKE ?"
		args = append(args, "%"+email+"%")
	}

	if search != "" {
		query += " AND (name LIKE ? OR email LIKE ?)"
		args = append(args, "%"+search+"%", "%"+search+"%")
	}

	query += " ORDER BY id LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func (r *UserRepository) FindByID(id int64) (*model.User, error) {
	var u model.User
	rows, err := r.db.Query(`SELECT id, name, email FROM users WHERE id = ?`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, errors.New("usuário não encontrado")
	}

	err = rows.Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// Update
func (r *UserRepository) Update(user *model.User) error {
	_, err := r.db.Exec(
		`UPDATE users SET name = ?, email = ? WHERE id = ?`,
		user.Name,
		user.Email,
		user.ID,
	)
	return err
}

// Delete
func (r *UserRepository) Delete(id int64) error {
	_, err := r.db.Exec(`DELETE FROM users WHERE id = ?`, id)
	return err
}
