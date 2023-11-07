package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/umizu/yomu/internal/data"
	"github.com/umizu/yomu/internal/events"
	"github.com/umizu/yomu/internal/models"
	"github.com/umizu/yomu/pkg/contracts"
)

type LibraryItemHandler struct {
	libraryItemStore data.LibraryItemStore
	bookStore        data.BookStore
}

func NewLibraryItemHandler(lStore data.LibraryItemStore, bStore data.BookStore) *LibraryItemHandler {
	return &LibraryItemHandler{
		libraryItemStore: lStore,
		bookStore:        bStore,
	}
}

func (h *LibraryItemHandler) LibraryItemGETHandler(c echo.Context) error {
	libraryItems, err := h.libraryItemStore.GetAll()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, models.ToResponse(libraryItems))
}

func (h *LibraryItemHandler) LibraryItemPUTHandler(c echo.Context) error {
	var req contracts.UpsertLibraryItemRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	book, err := h.bookStore.GetBookById(req.BookId)
	if err != nil {
		return err
	}
	if book == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "book not found"})
	}

	libraryItem := models.NewLibraryItemFromRequest(req)
	if err := h.libraryItemStore.Upsert(libraryItem); err != nil {
		return err
	}

	events.LibraryItemCh <- events.LibraryItemUpsertedEvent{Message: "todo: implement event"}
	return c.JSON(http.StatusCreated, libraryItem.ToResponse())
}
