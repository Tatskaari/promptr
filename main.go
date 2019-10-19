package main

import (
	"promptr/argprovider/unix"
	"promptr/prompt"
)

func main() {
	println(prompt.GeneratePrompt(new(unix.LinuxArgProvider)))
}
