package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s Server) RedirectToURL(c echo.Context) error{
	shortCode := c.Param("short_code")
	url, err := s.ShortenSvc.GetURL(shortCode)
	fmt.Println(shortCode)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}


	return c.Redirect(http.StatusTemporaryRedirect, url)
}