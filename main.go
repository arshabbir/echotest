package main

import (
	"echomod/api"

	"github.com/labstack/echo/v4"
)

func main() {
	apiHandler := api.NewServer()
	mux := echo.New()
	mux.GET("/user", apiHandler.HandleUser)
	mux.POST("/user", apiHandler.SaveUser)

	mux.Start(":8080")

}
