package prompt

import (
	"github.com/stretchr/testify/require"
	argsmock "promptr/prompt/promptargs/mock"
	"testing"
)

func TestGeneratePrompt(t *testing.T) {
	r := require.New(t)

	m := new(argsmock.Mock)

	m.On("Username").Return("user")
	m.On("Hostname").Return("host")
	m.On("WorkingDirectory").Return("/home/user/dev")
	m.On("Prompt").Return("$")

	result := GeneratePrompt(m)

	r.Equal("user@host:/home/user/dev $ ", result)

	m.AssertExpectations(t)

}