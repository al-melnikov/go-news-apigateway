package db

import (
	"comments/pkg/models"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	db *pgxpool.Pool
}

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

func (db *DB) AddComment(p models.Comment) (*uuid.UUID, error) {
	query := `INSERT INTO comments (news_id, parent_id, created_at, content) 
				VALUES ($1, $2, $3::timestamp, $4) RETURNING id;`

	var commentID uuid.UUID
	err := db.db.QueryRow(context.Background(), query, p.NewsID, p.ParentID, p.CreatedAt, p.Content).Scan(&commentID)
	return &commentID, err
}

func (db *DB) GetComments(newsID uuid.UUID) ([]models.Comment, error) {
	query := `SELECT id, news_id, parent_id, content, created_at::timestamp FROM comments
				WHERE news_id = $1 ORDER BY created_at DESC;`

	var comments []models.Comment

	rows, err := db.db.Query(context.Background(), query, newsID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var t models.Comment
		err = rows.Scan(
			&t.ID,
			&t.NewsID,
			&t.ParentID,
			&t.Content,
			&t.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		// добавление переменной в массив результатов
		comments = append(comments, t)

	}

	return comments, err
}
