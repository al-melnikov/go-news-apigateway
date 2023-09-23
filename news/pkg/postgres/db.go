package database

import (
	"context"
	"fmt"
	"news/pkg/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

const ItemsOnPage = 10

type DB struct {
	db *pgxpool.Pool
}

// Возвращает новый экземпляр базы данных
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

// Возвращает новость по ее ID и ошибку
func (db *DB) GetNewsByID(newsID uuid.UUID) (models.Post, error) {
	query := `SELECT id, title, content, created_at::timestamp, link FROM posts
				WHERE id = $1;`

	var post models.Post
	err := db.db.QueryRow(context.Background(), query, newsID).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
		&post.Link,
	)
	if err != nil {
		return models.Post{}, err
	}
	return post, nil
}

// Ищет статьи по регулярному выражению и возвращает массив из них,
// число страниц, текущую страницу и ошибку. Текущую страницу принимет как аргумент
// Если для нее нет постов, то возвращает для 1 страницы
func (db *DB) GetNewsByRegExp(regExp string, currentPage int) (
	posts []models.Post,
	pagesNum int,
	curPage int,
	err error,
) {

	queryCount := `SELECT COUNT(*) FROM posts
				WHERE title ~*$1;`

	var itemsNum int

	err = db.db.QueryRow(context.Background(), queryCount, regExp).Scan(&itemsNum)
	//fmt.Println(itemsNum)
	if err != nil {
		return nil, 0, 0, err
	}

	fmt.Println(currentPage)

	// Вернет результаты для 1 страницы если на запрашиваемой ничего нет
	if itemsNum <= (currentPage-1)*ItemsOnPage+1 {
		currentPage = 1
	}

	fmt.Println(currentPage)

	if itemsNum <= 0 {
		currentPage = 0
		//return nil, 0, 0, 0, errors.New("nothing found")
	}

	query := `SELECT id, title, content, created_at::timestamp, link FROM posts
				WHERE title ~*$1 ORDER BY created_at DESC LIMIT $2 OFFSET $3;`

	//var posts []models.Post

	//var currentPage int = 1

	rows, err := db.db.Query(context.Background(), query, regExp, ItemsOnPage, (currentPage-1)*ItemsOnPage)
	if err != nil {
		return nil, 0, 0, err
	}

	for rows.Next() {
		var t models.Post
		err = rows.Scan(
			&t.ID,
			&t.Title,
			&t.Content,
			&t.CreatedAt,
			&t.Link,
		)
		if err != nil {
			return nil, 0, 0, err
		}
		// добавление переменной в массив результатов
		posts = append(posts, t)

	}

	pagesNum = itemsNum / ItemsOnPage
	if itemsNum%ItemsOnPage > 0 {
		pagesNum++
	}

	return posts, pagesNum, currentPage, err
}
