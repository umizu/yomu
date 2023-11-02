package api

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/umizu/yomu/internal/data"
)

type APIServer struct {
	listenAddr string
	db         *sql.DB
	router     *echo.Echo
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
	}, nil
}

func (s *APIServer) Run() {
	s.router.HTTPErrorHandler = customHTTPErrorHandler

	s.RegisterBookRoutes(s.db)
	s.router.Logger.Fatal(s.router.Start(s.listenAddr))
}

func customHTTPErrorHandler(err error, c echo.Context) {
	c.Logger().Error(err)
	c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
}
