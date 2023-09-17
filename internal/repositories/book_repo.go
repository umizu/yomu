package repositories

import "github.com/umizu/yomu/internal/models"

type BookRepository interface {
	GetAll() ([]models.Book, error)
	GetByID(id string) (models.Book, error)
	Create(book models.Book) (models.Book, error)
}
