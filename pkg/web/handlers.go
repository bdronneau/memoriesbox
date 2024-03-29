package web

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

func (a *app) readyHandler(c echo.Context) error {
	err := a.repositories.PingDB()
	if err != nil {
		return NewSimpleResponse(c, http.StatusServiceUnavailable, "DB is not available")
	}

	return NewSimpleResponse(c, http.StatusOK, "I'm ready")
}

func (a *app) liveHandler(c echo.Context) error {
	return NewSimpleResponse(c, http.StatusOK, "I'm live good")
}

func (a *app) versionHandler(c echo.Context) error {
	version := os.Getenv("APP_VERSION")
	if version == "" {
		version = "development"
	}

	return c.String(http.StatusOK, version)
}

func (a *app) countMemories(c echo.Context) error {
	count, err := a.repositories.CountMemories()
	if err != nil {
		slog.Error("Can not count memories", "error", err)
		return c.JSON(http.StatusInternalServerError, "Oups check application log")
	}

	return c.String(http.StatusOK, fmt.Sprintf("%d", count))
}

func (a *app) getMemories(c echo.Context) error {
	memory, err := a.repositories.GetRandomMemories()
	if err != nil {
		slog.Error("Can not retrieve a random memory", "error", err)
		return c.JSON(http.StatusInternalServerError, "Oups check application log")
	}

	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"content": memory.Content,
		"date":    memory.Append,
		"author":  memory.Author,
	})
}

func (a *app) addMemory(c echo.Context) error {
	messages := []string{"miaou", "waou"}
	err := c.Render(http.StatusOK, "memory_add.html", map[string]interface{}{
		"messages": messages,
	})
	if err != nil {
		slog.Error("Can not render", "error", err)
	}

	return err
}

func (a *app) addAPIMemory(c echo.Context) error {
	var err error

	author := c.FormValue("author")
	if author == "" {
		err = formValidationErrors("No author", err)
	}

	dateRaw := c.FormValue("date")
	if dateRaw == "" {
		err = formValidationErrors("No date", err)
	}

	date, errTimeParse := time.Parse(time.DateOnly, dateRaw)
	if errTimeParse != nil {
		err = formValidationErrors("Date invalid format", err)
	}

	quote := c.FormValue("quote")
	if quote == "" {
		err = formValidationErrors("No quote", err)
	}

	if err != nil {
		return c.Render(http.StatusBadRequest, "memory_add.html", map[string]interface{}{
			"messages": err.Error(),
		})
	}

	err = a.repositories.AddMemory(quote, author, date)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("%v", err.Error()))
	}

	slog.Debug("Quote create", "author", author, "date", date, "quote", quote)

	return c.Redirect(301, "/")
}

func formValidationErrors(message string, err error) error {
	if err == nil {
		return errors.New(message)
	}

	return errors.Join(err, errors.New(message))
}
