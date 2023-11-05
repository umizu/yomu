package data

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/umizu/yomu/internal/types"
)

type BookStore interface {
	CreateBook(*types.Book) error
	GetBookById(id string) (*types.Book, error)
	GetAllBooks() ([]*types.Book, error)
}

type PostgresBookStore struct {
	db *sql.DB
}

func NewPostgresBookStore(db *sql.DB) *PostgresBookStore {
	return &PostgresBookStore{
		db: db,
	}
}

func (s *PostgresBookStore) CreateBook(book *types.Book) error {
	_, err := s.db.Exec(`
		INSERT INTO book (id, title, isbn, format, link, language)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, book.ID, book.Title, book.Isbn, book.Format, book.Link, book.Language)

	return err
}

func (s *PostgresBookStore) GetBookById(id string) (*types.Book, error) {
	row := s.db.QueryRow(`
		SELECT id, title, isbn, format, link, language
		FROM book
		WHERE id = $1
	`, id)

	var book types.Book

	if err := row.Scan(&book.ID, &book.Title, &book.Format, &book.Isbn, &book.Link, &book.Language); err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, nil
		default:
			return nil, err
		}
	}

	return &book, nil
}

func (s *PostgresBookStore) GetAllBooks() ([]*types.Book, error) {
	rows, err := s.db.Query(`
		SELECT id, title, isbn, format, link, language
		FROM book
	`)
	if err != nil {
		return nil, err
	}

	books := []*types.Book{}
	for rows.Next() {
		var book types.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Isbn, &book.Format, &book.Link, &book.Language); err != nil {
			return nil, err
		}
		books = append(books, &book)
	}
	return books, nil
}
