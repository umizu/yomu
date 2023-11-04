package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/umizu/yomu/internal/data"
)

type BookStatusHandler struct {
	bookStatusStore data.BookStatusStore
}

func NewBookStatusHandler(bookStatusStore data.BookStatusStore) *BookStatusHandler {
	return &BookStatusHandler{
		bookStatusStore: bookStatusStore,
	}
}

func (h *BookStatusHandler) BookStatusGETHandler(c echo.Context) error {
	bookStatus, err := h.bookStatusStore.GetAll()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, bookStatus)
}

func (h *BookStatusHandler) BookStatusPOSTHandler(c echo.Context) error {
	return c.JSON(http.StatusCreated, nil)
}
