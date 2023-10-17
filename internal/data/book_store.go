package data

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/umizu/yomu/internal/models"
)

type BookStore interface {
	CreateBook(*models.Book) error
	GetBookById(id string) (*models.Book, error)
	GetAllBooks() ([]*models.Book, error)
}

type PostgresBookStore struct {
	db *sql.DB
}

func NewPostgresBookStore(db *sql.DB) *PostgresBookStore {
	return &PostgresBookStore{
		db: db,
	}
}

func (s *PostgresBookStore) CreateBook(book *models.Book) error {
	_, err := s.db.Exec(`
		INSERT INTO book (id, title, mediaType, length)
		VALUES ($1, $2, $3, $4)
	`, book.Id, book.Title, book.MediaType, book.Length)

	return err
}

func (s *PostgresBookStore) GetBookById(id string) (*models.Book, error) {
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

func (s *PostgresBookStore) GetAllBooks() ([]*models.Book, error) {
	rows, err := s.db.Query(`
		SELECT id, title, mediaType, length
		FROM book
	`)

	if err != nil {
		return nil, err
	}

	books := []*models.Book{}

	for rows.Next() {
		var book models.Book

		if err := rows.Scan(&book.Id, &book.Title, &book.MediaType, &book.Length); err != nil {
			return nil, err
		}

		books = append(books, &book)
	}

	return books, nil
}
