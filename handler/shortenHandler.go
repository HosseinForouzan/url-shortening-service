package handler

import (
	"fmt"
	"net/http"

	"github.com/HosseinForouzan/url-shortening-service/shorten"
	"github.com/labstack/echo/v4"
)

func (s Server) CreateHandler(c echo.Context) error {
	var req shorten.ShortenRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp, err := s.ShortenSvc.CreateService(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, resp)
	
}

func (s Server) ReadHandler(c echo.Context) error {
	shortCode := c.Param("short_code")
	fmt.Println(shortCode)
	req := shorten.RetireveRequest{ShortCode: shortCode}

	resp, err := s.ShortenSvc.RetrieveService(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)

}