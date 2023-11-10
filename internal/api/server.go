package api

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/umizu/yomu/internal/data"
	"github.com/umizu/yomu/internal/events"
)

type APIServer struct {
	listenAddr string
	db         *sql.DB
	router     *echo.Echo
	msgch      chan interface{}
}

func NewAPIServer(listenAddr string) (*APIServer, error) {
	pgStore, err := data.NewPostgresStore()
	if err != nil {
		return nil, err
	}

	if err := pgStore.Init(); err != nil {
		return nil, err
	}

	return &APIServer{
		listenAddr: listenAddr,
		db:         pgStore.DB,
		router:     echo.New(),
		msgch:      make(chan interface{}),
	}, nil
}

func (s *APIServer) Run() {
	s.router.HTTPErrorHandler = customHTTPErrorHandler
	go events.Listen(s.msgch)
	s.RegisterRoutes()
	s.router.Logger.Fatal(s.router.Start(s.listenAddr))
}

func customHTTPErrorHandler(err error, c echo.Context) {
	c.Logger().Error(err)
	if err == echo.ErrNotFound {
		c.JSON(http.StatusNotFound, map[string]string{"error": "path not found"})
		return
	}
	c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
}
