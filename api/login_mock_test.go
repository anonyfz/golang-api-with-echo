package api_test

import "github.com/stretchr/testify/mock"

type MockLoginService struct {
	mock.Mock
}

func (m *MockLoginService) Login() error {
	args := m.Called()
	return args.Error(0)
}
