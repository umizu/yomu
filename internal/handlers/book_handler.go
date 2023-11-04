package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/umizu/yomu/internal/data"
	"github.com/umizu/yomu/internal/types"
	"github.com/umizu/yomu/pkg/contracts"
)

type BookHandler struct {
	bookStore data.BookStore
}

func NewBookHandler(bookStore data.BookStore) *BookHandler {
	return &BookHandler{
		bookStore: bookStore,
	}
}

func (h *BookHandler) BooksGETHandler(c echo.Context) error {
	books, err := h.bookStore.GetAllBooks()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, books)
}

func (h *BookHandler) BooksPOSTHandler(c echo.Context) error {
	var request contracts.CreateBookRequest
	if err := c.Bind(&request); err != nil {
		return err
	}
	if err := request.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	book := types.NewBookFromRequest(request)
	if err := h.bookStore.CreateBook(book); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, book)
}
