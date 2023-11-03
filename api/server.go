package api

import "github.com/labstack/echo/v4"

type server struct {
}

type Server interface {
	HandleUser(c echo.Context) error
	SaveUser(c echo.Context) error
}

func NewServer() Server {
	return &server{}
}
