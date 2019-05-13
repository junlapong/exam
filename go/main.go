package main

import (
	"net/http"
	"todo/task"

	echo "github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Todo API")
	})

	e.GET("/tasks", task.ListTask)
	e.GET("/tasks/:id", task.GetTask)
	e.POST("/tasks", task.AddTask)
	e.PUT("/tasks/:id", task.UpdateTask)
	e.DELETE("/tasks/:id", task.DeleteTask)

	e.Static("/static", "static")

	e.Logger.Fatal(e.Start(":1323"))
}
