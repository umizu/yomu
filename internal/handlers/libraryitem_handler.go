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
	liStore   data.LibraryItemStore
	bookStore data.BookStore
	msgch   chan interface{}
}

func NewLibraryItemHandler(lStore data.LibraryItemStore, bStore data.BookStore, msgch chan interface{}) *LibraryItemHandler {
	return &LibraryItemHandler{
		liStore:   lStore,
		bookStore: bStore,
		msgch:   msgch,
	}
}

func (h *LibraryItemHandler) LibraryItemGETHandler(c echo.Context) error {
	libraryItems, err := h.liStore.GetAll()
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
	if err := h.liStore.Upsert(libraryItem); err != nil {
		return err
	}

	h.msgch <- events.LibraryItemUpsertedEvent{Message: "libraryitem has been upserted!", Store: h.liStore}
	return c.JSON(http.StatusCreated, libraryItem.ToResponse())
}
