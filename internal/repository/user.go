package repository

import (
	"context"
	"extra_muse/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface{
	Save(model.NewUserData) error
	GetById(TgID int) (*model.User, error)
}

func CreateUserTable(ctx context.Context,  db *pgxpool.Pool) error {

	query := `
	CREATE TABLE IF NOT EXISTS users (
    tg_id INTEGER PRIMARY KEY, 
    username TEXT,
    first_name TEXT,
    last_name TEXT,
    balance REAL DEFAULT 0.0,  
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)`

		if _, err := db.Exec(ctx, query); err != nil {
			return err //контекст чтобы  Иметь возможность прервать операцию, если, например, запрос к БД выполняется слишком долго (завис) – через context.WithTimeout или context.WithCancel
} 
	return nil
}



