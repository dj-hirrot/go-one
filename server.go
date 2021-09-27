package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", func(c echo.Context) error { return hello(c) })

	e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) (err error) {
	c.JSON(200, "users2")
	return
}
