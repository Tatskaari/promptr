package prompt

import (
	"fmt"
	"promptr/prompt/promptargs"
)


func GeneratePrompt(args promptargs.PromptArgs) string {
	return fmt.Sprintf("%s@%s:%s %s ", args.Username(), args.Hostname(), args.WorkingDirectory(), args.Prompt())
}