package data

import (
	"database/sql"

	"github.com/umizu/yomu/internal/types"
)

type LibraryItemStore interface {
	Create(*types.LibraryItem) error
	GetAll() ([]*types.LibraryItem, error)
}

type PostgresLibraryItemStore struct {
	db *sql.DB
}

func NewPostgresLibraryItemStore(db *sql.DB) *PostgresLibraryItemStore {
	return &PostgresLibraryItemStore{db: db}
}

func (s *PostgresLibraryItemStore) Create(bs *types.LibraryItem) error {
	_, err := s.db.Exec("INSERT INTO library_item (id, book_id, status) VALUES ($1, $2, $3)", bs.Id, bs.BookId, bs.Status)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresLibraryItemStore) GetAll() ([]*types.LibraryItem, error) {
	rows, err := s.db.Query("SELECT id, book_id, status FROM library_item")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	libraryItems := []*types.LibraryItem{}
	for rows.Next() {
		bs := &types.LibraryItem{}
		if err := rows.Scan(&bs.Id, &bs.BookId, &bs.Status); err != nil {
			return nil, err
		}
		libraryItems = append(libraryItems, bs)
	}
	return libraryItems, nil
}
