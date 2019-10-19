package prompt

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type argsMock struct {
	username string
	hostname string
	workingDir string
	prompt string
}

func (a *argsMock) Username() string {
	return a.username
}

func (a *argsMock) Hostname() string {
	return a.hostname
}

func (a *argsMock) Prompt() string {
	return a.workingDir
}

func (a *argsMock) WorkingDirectory() string {
	return a.prompt
}

func TestGeneratePrompt(t *testing.T) {
	r := require.New(t)

	m := &argsMock{
		username:   "user",
		hostname:   "host",
		workingDir: "/home/user/dev",
		prompt:     "$",
	}

	r.Equal("user@host:/home/user/dev $ ",  GeneratePrompt(m))
}