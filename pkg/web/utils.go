package web

import "github.com/labstack/echo/v4"

// HTTPSuccess struct for simple response
type HTTPSuccess struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"I'm a description"`
}

// HTTPError struct for formating all erors
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// NewError generator
func NewError(c echo.Context, status int, err error) error {
	er := &HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	return c.JSON(status, er)
}

// NewSimpleResponse generator
func NewSimpleResponse(c echo.Context, status int, message string) error {
	er := &HTTPSuccess{
		Code:    status,
		Message: message,
	}
	return c.JSON(status, er)
}
