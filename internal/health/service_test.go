package health

import (
	"bytes"
	"errors"
	"log/slog"
	"testing"
)

func TestService_Check_Success(t *testing.T) {
	mockLogger := slog.New(slog.NewTextHandler(&bytes.Buffer{}, nil))

	mockRepo := new(MockHealthRepository)
	mockRepo.On("Select").Return(nil)

	service := NewService()
	err := service.Check(mockRepo, mockLogger)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	mockRepo.AssertExpectations(t)
}

func TestService_Check_Failure(t *testing.T) {
	mockLogger := slog.New(slog.NewTextHandler(&bytes.Buffer{}, nil))

	mockRepo := new(MockHealthRepository)
	mockRepo.On("Select").Return(errors.New("data fetch error"))

	service := NewService()
	err := service.Check(mockRepo, mockLogger)
	if err == nil {
		t.Errorf("expected an error, got nil")
	}

	if err.Error() != "data fetch error" {
		t.Errorf("expected error message 'data fetch error', got '%v'", err.Error())
	}

	mockRepo.AssertExpectations(t)
}
