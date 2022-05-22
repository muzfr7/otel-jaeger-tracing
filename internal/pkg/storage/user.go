package storage

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/muzfr7/otel-jaeger-tracing/internal/pkg/trace"
)

type UserStorer interface {
	Insert(ctx context.Context, user User) error
}

type User struct {
	ID   int
	Name string
}

var _ UserStorer = UserStorage{}

type UserStorage struct {
	database *sql.DB
}

func NewUserStorage(dtb *sql.DB) UserStorage {
	return UserStorage{
		database: dtb,
	}
}

func (u UserStorage) Insert(ctx context.Context, user User) error {
	// Create a child span.
	ctx, span := trace.NewSpan(ctx, "UserStorage.Insert", nil)
	defer span.End()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	if _, err := u.database.ExecContext(ctx, `INSERT INTO users (name) VALUES (?)`, user.Name); err != nil {
		log.Println(err)

		return fmt.Errorf("insert: failed to execute query")
	}

	return nil
}
