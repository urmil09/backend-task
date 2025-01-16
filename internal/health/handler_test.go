package health

import (
	"bytes"
	"errors"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogHandler_Success(t *testing.T) {
	mockLogger := slog.New(slog.NewTextHandler(&bytes.Buffer{}, nil))

	mockRepo := new(MockHealthRepository)
	mockRepo.On("Select").Return(nil)

	handler := NewHandler()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()
	handler.CheckHealth(mockRepo, mockLogger, w, req)

	res := w.Result()
	defer res.Body.Close()
	expectedStatus := http.StatusOK

	if res.StatusCode != expectedStatus {
		t.Errorf("expected status %d, got %d", expectedStatus, res.StatusCode)
	}

	mockRepo.AssertExpectations(t)
}

func TestCheckHealth_Failure(t *testing.T) {
	mockLogger := slog.New(slog.NewTextHandler(&bytes.Buffer{}, nil))

	mockRepo := new(MockHealthRepository)
	mockRepo.On("Select").Return(errors.New("health check failed"))

	handler := NewHandler()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	handler.CheckHealth(mockRepo, mockLogger, w, req)
	expectedStatus := http.StatusInternalServerError

	res := w.Result()
	if res.StatusCode != expectedStatus {
		t.Errorf("expected status %d, got %d", expectedStatus, res.StatusCode)
	}

	mockRepo.AssertExpectations(t)
}
