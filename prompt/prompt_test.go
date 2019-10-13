package prompt

import (
	"github.com/stretchr/testify/require"
	"promptr/prompt/mockpromptargs"
	"testing"
)

func TestGeneratePrompt(t *testing.T) {
	r := require.New(t)

	mock := new(mockpromptargs.Mock)

	mock.On("Username").Return("user")
	mock.On("Hostname").Return("host")
	mock.On("WorkingDirectory").Return("/home/user/dev")
	mock.On("Prompt").Return("$")

	result := GeneratePrompt(mock)

	r.Equal("user@host:/home/user/dev $ ", result)

	mock.AssertExpectations(t)

}