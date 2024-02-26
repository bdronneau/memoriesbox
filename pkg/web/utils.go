package web

import "github.com/labstack/echo/v4"

// HTTPSuccess struct for simple response
type HTTPSuccess struct {
	Message string `json:"message" example:"I'm a description"`
	Code    int    `json:"code" example:"200"`
}

// HTTPError struct for formating all erors
type HTTPError struct {
	Message string `json:"message" example:"status bad request"`
	Code    int    `json:"code" example:"400"`
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
