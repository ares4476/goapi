package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos = map[int]*todo{
	1: &todo{ID: 1, Title: "pay phone bill1", Status: "active"},
}

func todoHandler(c echo.Context) error {
	items := []*todo{}
	for _, item := range todos {
		items = append(items, item)
	}

	return c.JSON(http.StatusOK, items)
}

func createTodoHandler(e echo.Context) error {
	t := todo{}
	if err := e.Bind(&t); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	id := len(todos)
	id++
	t.ID = id
	todos[t.ID] = &t
	return e.JSON(http.StatusCreated, "create todo")
}
func getTodoHandlerByid(c echo.Context) error {
	var id int
	err := echo.PathParamsBinder(c).Int("id", &id).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	t, ok := todos[id]
	if !ok {
		return c.JSON(http.StatusOK, map[int]string{})
	}
	return c.JSON(http.StatusOK, t)

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
	e.GET("/todos/:id", getTodoHandlerByid)
	e.POST("/todos", createTodoHandler)

	port := os.Getenv("PORT")
	log.Println("port:", port)
	e.Start(":" + port)
}
