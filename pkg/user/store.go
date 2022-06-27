package user

import (
	"context"
	"fmt"
	"log"
	"msqrd/pkg/model"
	"os"

	"github.com/jackc/pgx/v4"
)

type Store struct {
	db *pgx.Conn
}

func (s *Store) New() *Store {
	return &Store{}
}

func (s *Store) Open(ctx context.Context) error {
	conn, err := pgx.Connect(ctx, os.Getenv("DB_URL"))
	if err != nil {
		return nil
	}

	log.Println("success db connection")
	s.db = conn

	return nil
}

func (s *Store) CreateUser(ctx context.Context, u *model.User) (*model.User, error) {
	if err := s.db.QueryRow(
		ctx,
		"INSERT INTO users (username, email, encrypted_password) VALUES ($1, $2, $3) RETURNING id",
		u.Username, u.Email, u.EncryptedPassword,
	).Scan(&u.Id); err != nil {
		return nil, err
	}

	return u, nil
}

func (s *Store) DeleteUser(ctx context.Context, id int32) (string, error) {
	_, err := s.db.Query(
		ctx,
		"DELETE FROM users WHERE id = $1",
		id,
	)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("User with id: %d was succesfully deleted", id), nil
}
