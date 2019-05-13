package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"
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

	// Start server
	go func() {
		if err := e.Start(":1323"); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	// with a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
