package main

import (
	"github.com/jessevdk/go-flags"
	"log"
	"promptr/prompt"
)

var options struct {
	Username string `short:"u"`
	Hostname string `short:"n"`
	Prompt string `short:"p"`
	WorkingDirectory string `short:"d"`
}

//TODO(jpoole): Make this retrieve these values from the command line or env vars
type opts struct {}

func (o *opts) Username() string {
	return options.Username
}

func (o *opts) Hostname() string {
	return options.Hostname
}

func (o *opts) Prompt() string {
	return options.Prompt
}

func (o *opts) WorkingDirectory() string {
	return options.WorkingDirectory
}

func main() {
	var parser = flags.NewParser(&options, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		log.Fatalf("failed to parse flags: %v", err)
	}

	println(prompt.GeneratePrompt(new(opts)))
}
