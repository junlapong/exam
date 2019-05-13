package task

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestTask(t *testing.T) {

	t.Helper()

	// Setup
	e := echo.New()

	t.Run("list tasks", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/tasks")

		// Assertions
		if err := ListTask(c); err != nil {
		}
	})

	t.Run("get task", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/tasks/:id")
		c.SetParamNames("1")

		// Assertions
		if err := GetTask(c); err != nil {
		}
	})

	t.Run("add task", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{}"))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/tasks")

		// Assertions
		if err := AddTask(c); err != nil {
		}
	})

	t.Run("update task", func(t *testing.T) {

	})

	t.Run("delete task", func(t *testing.T) {

	})
	 
}
