package mock

import (
	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) Username() string {
	args := m.Called()
	return args.String(0)
}

func (m *Mock) Hostname() string {
	args := m.Called()
	return args.String(0)
}

func (m *Mock) Prompt() string {
	args := m.Called()
	return args.String(0)
}

func (m *Mock) WorkingDirectory() string {
	args := m.Called()
	return args.String(0)
}



