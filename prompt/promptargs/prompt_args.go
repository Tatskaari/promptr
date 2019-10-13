package promptargs

type PromptArgs interface {
	Username() string
	Hostname() string
	Prompt() string
	WorkingDirectory() string
}