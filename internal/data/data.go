package data

import "database/sql"

const (
	connStr = "postgres://postgres:yomu@localhost:7000?sslmode=disable"
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
	_, err := s.DB.Exec(`
		CREATE TABLE IF NOT EXISTS book (
			id UUID PRIMARY KEY,
			title TEXT,
			mediaType TEXT,
			length INT)
	`)
	return err
}
