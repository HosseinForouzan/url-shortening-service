package handler

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/HosseinForouzan/url-shortening-service/shorten"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	ShortenSvc shorten.Shorten
}

func New(shortenSvc shorten.Shorten) Server {
	return Server{ShortenSvc: shortenSvc }

}



func (s Server) SetRoutes() {
	e := echo.New()

  // Middleware
	e.Use(middleware.RequestLogger()) // use the default RequestLogger middleware with slog logger
	e.Use(middleware.Recover()) // recover panics as errors for proper error handling

	// Routes
	
	e.GET("/", hello)
	
	shorten := e.Group("/shorten")

	shorten.POST("/create", s.CreateHandler)

	// Start server
	if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)

	}
}

func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}