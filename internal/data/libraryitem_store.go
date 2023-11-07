package data

import (
	"database/sql"

	"github.com/umizu/yomu/internal/models"
)

type LibraryItemStore interface {
	Upsert(*models.LibraryItem) error
	GetAll() ([]*models.LibraryItem, error)
	GetByBookId(string) (*models.LibraryItem, error)
}

type PostgresLibraryItemStore struct {
	db *sql.DB
}

func NewPostgresLibraryItemStore(db *sql.DB) *PostgresLibraryItemStore {
	return &PostgresLibraryItemStore{db: db}
}

func (s *PostgresLibraryItemStore) Upsert(item *models.LibraryItem) error {
	existingItem, err := s.GetByBookId(item.BookId)
	if err != nil {
		return err
	}
	if existingItem == nil {
		_, err = s.db.Exec("INSERT INTO library_item (id, book_id, status) VALUES ($1, $2, $3)", item.Id, item.BookId, item.Status)
		if err != nil {
			return err
		}
		return nil
	}

	_, err = s.db.Exec("UPDATE library_item SET status = $1 WHERE book_id = $2", item.Status, item.BookId)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresLibraryItemStore) GetAll() ([]*models.LibraryItem, error) {
	rows, err := s.db.Query("SELECT id, book_id, status FROM library_item")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	libraryItems := []*models.LibraryItem{}
	for rows.Next() {
		items := &models.LibraryItem{}
		if err := rows.Scan(&items.Id, &items.BookId, &items.Status); err != nil {
			return nil, err
		}
		libraryItems = append(libraryItems, items)
	}
	return libraryItems, nil
}

func (s *PostgresLibraryItemStore) GetByBookId(bookId string) (*models.LibraryItem, error) {
	row := s.db.QueryRow("SELECT id, book_id, status FROM library_item WHERE book_id = $1", bookId)
	item := &models.LibraryItem{}
	err := row.Scan(&item.Id, &item.BookId, &item.Status)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return item, nil
}
