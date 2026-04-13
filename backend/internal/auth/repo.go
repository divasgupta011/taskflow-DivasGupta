package auth

import (
	"context"
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(ctx context.Context, u *User) error {
	query := `
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`

	return r.db.QueryRowContext(ctx, query,
		u.Name, u.Email, u.Password,
	).Scan(&u.ID, &u.CreatedAt)
}

func (r *Repository) GetByEmail(ctx context.Context, email string) (*User, error) {
	query := `
		SELECT id, name, email, password, created_at
		FROM users
		WHERE email = $1
	`

	var u User
	err := r.db.QueryRowContext(ctx, query, email).
		Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &u, nil
}
