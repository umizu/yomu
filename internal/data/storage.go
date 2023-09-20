package data

import (
	"database/sql"
	_ "github.com/lib/pq"

	"github.com/umizu/yomu/internal/models"
)

type Store interface {
	CreateBook(*models.Book) error
	GetBookById(id string) (*models.Book, error)
	GetAllBooks() ([]*models.Book, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "postgres://postgres:yomu@localhost:7000?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) CreateBook(book *models.Book) error {
	_, err := s.db.Exec(`
		INSERT INTO book (id, title, mediaType, length)
		VALUES ($1, $2, $3, $4)
	`, book.Id, book.Title, book.MediaType, book.Length)

	return err
}

func (s *PostgresStore) GetBookById(id string) (*models.Book, error) {
	row := s.db.QueryRow(`
		SELECT id, title, mediaType, length
		FROM book
		WHERE id = $1
	`, id)

	var book models.Book

	if err := row.Scan(&book.Id, &book.Title, &book.MediaType, &book.Length); err != nil {
		return nil, err
	}

	return &book, nil
}

func (s *PostgresStore) GetAllBooks() ([]*models.Book, error) {
	rows, err := s.db.Query(`
		SELECT id, title, mediaType, length
		FROM book
	`)

	if err != nil {
		return nil, err
	}

	var books []*models.Book

	for rows.Next() {
		var book models.Book

		if err := rows.Scan(&book.Id, &book.Title, &book.MediaType, &book.Length); err != nil {
			return nil, err
		}

		books = append(books, &book)
	}

	return books, nil
}
