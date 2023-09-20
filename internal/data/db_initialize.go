package data

func (s *PostgresStore) Init() error {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS book (
			id UUID PRIMARY KEY,
			title TEXT,
			mediaType TEXT,
			length INT)
	`)

	return err
}
