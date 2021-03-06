package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type todo struct {
	ID     int    `JSON:"id"`
	Title  string `JSON:"title"`
	Status string `JSON:"status"`
}

var totos = map[int]*todo{
	1: &todo{ID: 1, Title: "pay phone bill", Status: "active"},
}

func todoHandler(c echo.Context) error {
	items := []*todo{}
	for _, item := range totos {
		items = append(items, item)
	}

	return c.JSON(http.StatusOK, items)
}

func helloHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "watchrapon",
	})
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/hello", helloHandler)
	e.GET("/todos", todoHandler)
	port := os.Getenv("PORT")
	log.Println("port:", port)
	e.Start(":" + port)
}
