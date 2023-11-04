package data

import (
	"database/sql"

	"github.com/umizu/yomu/internal/types"
)

type BookStatusStore interface {
	Create(*types.BookStatus) error
	GetAll() ([]*types.BookStatus, error)
}

type PostgresBookStatusStore struct {
	db *sql.DB
}

func NewPostgresBookStatusStore(db *sql.DB) *PostgresBookStatusStore {
	return &PostgresBookStatusStore{db: db}
}

func (s *PostgresBookStatusStore) Create(bs *types.BookStatus) error {
	_, err := s.db.Exec("INSERT INTO book_status (id, book_id, status) VALUES ($1, $2, $3)", bs.Id, bs.BookId, bs.Status)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresBookStatusStore) GetAll() ([]*types.BookStatus, error) {
	rows, err := s.db.Query("SELECT id, book_id, status FROM book_status")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookStatus []*types.BookStatus
	for rows.Next() {
		bs := &types.BookStatus{}
		if err := rows.Scan(&bs.Id, &bs.BookId, &bs.Status); err != nil {
			return nil, err
		}
		bookStatus = append(bookStatus, bs)
	}
	return bookStatus, nil
}
