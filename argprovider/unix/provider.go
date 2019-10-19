package unix

import (
	"os"
	"os/exec"
	"strings"
)

type LinuxArgProvider struct {}

func (o *LinuxArgProvider) Username() string {
	return os.Getenv("USER")
}

func (o *LinuxArgProvider) Hostname() string {
	cmd :=  exec.Command("hostname")

	b, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	out := string(b)
	return strings.TrimSuffix(out, "\n")
}

func (o *LinuxArgProvider) Prompt() string {
	return "$"
}

func (o *LinuxArgProvider) WorkingDirectory() string {
	return os.Getenv("PWD")
}