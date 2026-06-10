package repository

import (
	"context"
	"extra_muse/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateGenerationsTablle(ctx context.Context, db *pgxpool.Pool) error {
	query := `
	CREATE TABLE IF NOT EXISTS generations (
			id SERIAL PRIMARY KEY,
			tg_id BIGINT NOT NULL,
			prompt TEXT,
			status VARCHAR(20) NOT NULL DEFAULT 'done',
			gen_id VARCHAR(100),
			track_url1 TEXT,
			image_url1 TEXT,
			title1 TEXT,
			track_url2 TEXT,
			image_url2 TEXT,
			title2 TEXT,
			error TEXT
	)`

	if _, err := db.Exec(ctx, query); err != nil {
return err //контекст чтобы  Иметь возможность прервать операцию, если, например, запрос к БД выполняется слишком долго (завис) – через context.WithTimeout или context.WithCancel
	} 

	return nil
}

type GenerationRepository interface{
	Save() error
	GetUserAll() ([]model.Track, error)
}