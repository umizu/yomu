package data

import "database/sql"

const (
	connStr = "postgres://postgres:yomu@localhost:9000?sslmode=disable"
)

type Store interface {
	Init() error
}

type PostgresStore struct {
	DB *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		DB: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	if _, err := s.DB.Exec(`
		CREATE TABLE IF NOT EXISTS book (
			id UUID PRIMARY KEY,
			title TEXT NOT NULL,
			isbn TEXT,
			format TEXT,
			link TEXT,
			language TEXT)
	`); err != nil {
		return err
	}

	if _, err := s.DB.Exec(`
		CREATE TABLE IF NOT EXISTS library_item (
			id UUID PRIMARY KEY,
			book_id UUID REFERENCES book(id) NOT NULL,
			status INT NOT NULL)
	`); err != nil {
		return err
	}
	
	if _, err := s.DB.Exec(`
		CREATE TABLE IF NOT EXISTS activity (
			id UUID PRIMARY KEY,
			book_id UUID REFERENCES book(id) NOT NULL,
			status INT NOT NULL)
	`); err != nil {
		return err
	}
	return nil
}
