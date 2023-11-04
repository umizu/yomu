package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/umizu/yomu/internal/data"
)

type LibraryItemHandler struct {
	libraryItemStore data.LibraryItemStore
}

func NewLibraryItemHandler(libraryItemStore data.LibraryItemStore) *LibraryItemHandler {
	return &LibraryItemHandler{
		libraryItemStore: libraryItemStore,
	}
}

func (h *LibraryItemHandler) LibraryItemGETHandler(c echo.Context) error {
	libraryItems, err := h.libraryItemStore.GetAll()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, libraryItems)
}

func (h *LibraryItemHandler) LibraryItemPOSTHandler(c echo.Context) error {
	return c.JSON(http.StatusCreated, nil)
}
