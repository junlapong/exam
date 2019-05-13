package task

import (
	"net/http"
	"strconv"

	"log"

	echo "github.com/labstack/echo/v4"
)

var seq = 1

// Task data
type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func ListTask(c echo.Context) error {
	tasks := make([]Task, 3)
	tasks[0] = Task{1, "Jun", "XXX"}
	tasks[1] = Task{2, "Jul", "YYY"}
	tasks[2] = Task{3, "Aug", "ZZZ"}

	return c.JSON(http.StatusOK, tasks)
}

func AddTask(c echo.Context) error {
	t := new(Task)
	if err := c.Bind(t); err != nil {
		return err
	}

	t.ID = seq
	seq++

	log.Printf("ADD: %d\n", t.ID)

	return c.JSON(http.StatusCreated, t)
}

func GetTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	log.Printf("GET: %d\n", id)

	t := new(Task)
	t.ID = id
	t.Name = "Jan"
	t.Desc = "AAA"

	return c.JSON(http.StatusOK, t)
}

func UpdateTask(c echo.Context) error {
	t := new(Task)
	if err := c.Bind(t); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))
	t.ID = id
	log.Printf("UPDATE: %d\n", id)

	return c.JSON(http.StatusOK, t)
}

func DeleteTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	log.Printf("DELETE: %d\n", id)

	return c.NoContent(http.StatusNoContent)
}
