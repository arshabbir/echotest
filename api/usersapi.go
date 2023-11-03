package api

import (
	"echomod/model"
	"echomod/utils"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *server) HandleUser(c echo.Context) error {

	log.Println("User handler invoked")
	usr := c.QueryParam("name")
	return c.JSON(http.StatusOK, model.User{Name: usr, City: "Hyd/Bang"})
}

func (s *server) SaveUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, &utils.ApiError{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusAccepted, user)

}
