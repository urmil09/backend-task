package health

import "github.com/stretchr/testify/mock"

type MockHealthRepository struct {
	mock.Mock
}

func (m *MockHealthRepository) Select() error {
	args := m.Called()
	return args.Error(0)
}
