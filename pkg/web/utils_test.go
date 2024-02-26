package web

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestNewSimpleResponse(t *testing.T) {
	// create a new echo instance
	e := echo.New()
	// create a new http request
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	// create a new http response recorder
	rec := httptest.NewRecorder()

	_ = NewSimpleResponse(e.NewContext(req, rec), http.StatusOK, "test message")

	assert.True(t, json.Valid(rec.Body.Bytes()))
	assert.Contains(t, rec.Body.String(), `"code":200`)
	assert.Contains(t, rec.Body.String(), `"message":"test message"`)
}

func TestNewError(t *testing.T) {
	// create a new echo instance
	e := echo.New()
	// create a new http request
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	// create a new http response recorder
	rec := httptest.NewRecorder()

	err := errors.New("test error")
	_ = NewError(e.NewContext(req, rec), http.StatusInternalServerError, err)

	expectedCode := http.StatusInternalServerError
	if expectedCode != rec.Code {
		t.Errorf("expected %d but got %d", expectedCode, rec.Code)
	}

	assert.True(t, json.Valid(rec.Body.Bytes()))
	assert.Contains(t, rec.Body.String(), `"code":500`)
	assert.Contains(t, rec.Body.String(), `"message":"test error"`)
}
