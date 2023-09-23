package db

import (
	"context"
	"rss-agg/pkg/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	db *pgxpool.Pool
}

// Возвращает экземпляр базы данных
func New(constr string) (*DB, error) {
	db, err := pgxpool.New(context.Background(), constr)
	if err != nil {
		return nil, err
	}
	s := DB{
		db: db,
	}
	return &s, nil
}

// Добавляет в базу данных пост
func (db *DB) AddPost(p models.Post) error {
	query := `INSERT INTO posts (title, content, link, created_at) 
				VALUES ($1, $2, $3, $4);`

	err := db.db.QueryRow(context.Background(), query, p.Title, p.Content, p.Link, p.CreatedAt).Scan()
	return err

}
