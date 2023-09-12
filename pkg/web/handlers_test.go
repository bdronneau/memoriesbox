package web

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/bdronneau/memoriesbox/pkg/logger"
	"github.com/bdronneau/memoriesbox/pkg/mocks"
	"github.com/bdronneau/memoriesbox/pkg/repositories"
	"github.com/bdronneau/memoriesbox/pkg/repositories/models"
	"go.uber.org/mock/gomock"

	"go.uber.org/zap/zaptest"
)

const EXPECT_STATUS_CODE = "expected status code %d but got %d"
const EXPECT_CONTENT = "expected response body to contain %q but got %q"

func bootstrapWebApp(t *testing.T, repoApp repositories.App) App {
	var (
		address       = "localhost"
		debug         = false
		port          = 1080
		featAddMemory = true
	)
	webConfig := Config{
		address:       &address,
		debug:         &debug,
		port:          &port,
		featAddMemory: &featAddMemory,
	}

	return New(webConfig, os.DirFS("../../cmd/webapp"), logger.App{Sugar: zaptest.NewLogger(t).Sugar()}, repoApp)
}

func TestGetRandomMemories(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoApp := mocks.NewRepositories(ctrl)

	expectedMemory := models.Memory{
		ID:      1,
		Author:  "John Doe",
		Content: "This is a memory",
		Append:  "2022-01-01",
	}
	repoApp.EXPECT().GetRandomMemories().Return(expectedMemory, nil)

	webApp := bootstrapWebApp(t, repoApp)
	// create a new http request
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	// create a new http response recorder
	rec := httptest.NewRecorder()
	// handle the request with the app's echo instance
	webApp.GetEcho().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf(EXPECT_STATUS_CODE, http.StatusOK, rec.Code)
	}

	expectedAuthor := "John Doe"
	expectedAppend := "2022-01-01"
	expectedContent := "This is a memory"
	if !strings.Contains(rec.Body.String(), expectedAuthor) || !strings.Contains(rec.Body.String(), expectedAppend) || !strings.Contains(rec.Body.String(), expectedContent) {
		t.Errorf("expected response body to contain %q, %q and %q but got %q", expectedAuthor, expectedAppend, expectedContent, rec.Body.String())
	}
}

func TestLive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoApp := mocks.NewRepositories(ctrl)

	webApp := bootstrapWebApp(t, repoApp)
	// create a new http request
	req := httptest.NewRequest(http.MethodGet, "/probes/live", nil)
	// create a new http response recorder
	rec := httptest.NewRecorder()
	// handle the request with the app's echo instance
	webApp.GetEcho().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf(EXPECT_STATUS_CODE, http.StatusOK, rec.Code)
	}

	expected := "I'm live good"
	if !strings.Contains(rec.Body.String(), expected) {
		t.Errorf(EXPECT_CONTENT, expected, rec.Body.String())
	}
}

func TestCountMemories(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoApp := mocks.NewRepositories(ctrl)

	var count int64 = 2
	repoApp.EXPECT().CountMemories().Return(count)

	webApp := bootstrapWebApp(t, repoApp)
	// create a new http request
	req := httptest.NewRequest(http.MethodGet, "/api/memories/count", nil)
	// create a new http response recorder
	rec := httptest.NewRecorder()
	// handle the request with the app's echo instance
	webApp.GetEcho().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf(EXPECT_STATUS_CODE, http.StatusOK, rec.Code)
	}

	expected := "2"
	if !strings.Contains(rec.Body.String(), expected) {
		t.Errorf(EXPECT_CONTENT, expected, rec.Body.String())
	}
}

func TestVersionHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoApp := mocks.NewRepositories(ctrl)

	webApp := bootstrapWebApp(t, repoApp)
	// create a new http request
	req := httptest.NewRequest(http.MethodGet, "/version", nil)
	// create a new http response recorder
	rec := httptest.NewRecorder()
	// handle the request with the app's echo instance
	webApp.GetEcho().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf(EXPECT_STATUS_CODE, http.StatusOK, rec.Code)
	}

	expected := "development"
	if !strings.Contains(rec.Body.String(), expected) {
		t.Errorf(EXPECT_CONTENT, expected, rec.Body.String())
	}
}

func TestFormValidationErrors(t *testing.T) {
	testCases := []struct {
		name        string
		message     string
		err         error
		expectedErr error
	}{
		{
			name:        "NoError",
			message:     "Validation failed",
			err:         nil,
			expectedErr: errors.New("Validation failed"),
		},
		{
			name:        "WithError",
			message:     "Validation failed",
			err:         errors.New("Specific error"),
			expectedErr: errors.Join(errors.New("Specific error"), errors.New("Validation failed")),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := formValidationErrors(tc.message, tc.err)
			if result == nil && tc.expectedErr != nil {
				t.Errorf("Expected an error, but got nil")
			} else if result != nil && tc.expectedErr == nil {
				t.Errorf("Expected no error, but got: %v", result)
			} else if result != nil && tc.expectedErr != nil && result.Error() != tc.expectedErr.Error() {
				t.Errorf("Expected error message '%s', but got '%s'", tc.expectedErr.Error(), result.Error())
			}
		})
	}
}
