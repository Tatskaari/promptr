package prompt

import "fmt"

type PromptArgs interface {
	Username() string
	Hostname() string
	Prompt() string
	WorkingDirectory() string
}

func GeneratePrompt(args PromptArgs) string {
	return fmt.Sprintf("%s@%s:%s %s ", args.Username(), args.Hostname(), args.WorkingDirectory(), args.Prompt())
}