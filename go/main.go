package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// Task data
type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

var seq = 1

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Todo API")
	})

	e.GET("/tasks", listTask)
	e.GET("/tasks/:id", getTask)
	e.POST("/tasks", addTask)
	e.PUT("/tasks/:id", updateTask)
	e.DELETE("/tasks/:id", deleteTask)

	e.Static("/static", "static")

	e.Logger.Fatal(e.Start(":1323"))
}

//----------
// Handlers
//----------

func listTask(c echo.Context) error {
	tasks := make([]Task, 3)
	tasks[0] = Task{1, "Jun", "XXX"}
	tasks[1] = Task{2, "Jul", "YYY"}
	tasks[2] = Task{3, "Aug", "ZZZ"}

	return c.JSON(http.StatusOK, tasks)
}

func addTask(c echo.Context) error {
	t := new(Task)
	if err := c.Bind(t); err != nil {
		return err
	}

	t.ID = seq
	seq++

	return c.JSON(http.StatusCreated, t)
}

func getTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Printf("GET: %d\n", id)

	t := new(Task)
	t.ID = id
	t.Name = "Jan"
	t.Desc = "AAA"

	return c.JSON(http.StatusOK, t)
}

func updateTask(c echo.Context) error {
	t := new(Task)
	if err := c.Bind(t); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))
	t.ID = id
	fmt.Printf("UPDATE: %d\n", id)

	return c.JSON(http.StatusOK, t)
}

func deleteTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Printf("DELETE: %d\n", id)

	return c.NoContent(http.StatusNoContent)
}
