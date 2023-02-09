package web

import (
	"fmt"
	"net/http"
	"os"

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
	return c.String(http.StatusOK, fmt.Sprintf("%d", a.repositories.CountMemories()))
}

func (a *app) getMemories(c echo.Context) error {

	memory, err := a.repositories.GetRandomMemories()
	if err != nil {
		a.logger.Errorf("Can not retrieve a random memory %v", err)
		return c.JSON(http.StatusInternalServerError, "Oups check application log")
	}

	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"content": memory.Content,
		"date":    memory.Append,
		"author":  memory.Author,
	})
}
