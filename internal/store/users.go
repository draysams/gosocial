package store

import (
	"context"
	"database/sql"
)

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	RoleID    int    `json:"role_id"`
	CreatedAt string `json:"created_at"`
}

type UsersStore struct {
	db *sql.DB
}

func (s *UsersStore) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users (username, email, password, role_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at
	`
	row := s.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password, user.RoleID)

	err := row.Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
