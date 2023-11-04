package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/umizu/yomu/internal/data"
	"github.com/umizu/yomu/internal/types"
	"github.com/umizu/yomu/pkg/contracts"
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
	return c.JSON(http.StatusOK, types.ToResponse(libraryItems))
}

func (h *LibraryItemHandler) LibraryItemPOSTHandler(c echo.Context) error {
	var req contracts.CreateLibraryItemRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	libraryItem := types.NewLibraryItemFromRequest(req)
	if err := h.libraryItemStore.Create(libraryItem); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, libraryItem.ToResponse())
}
