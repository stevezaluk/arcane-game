package prompt

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/stevezaluk/arcane-game/game"
	"os"
)

/*
CommandLinePrompt - Simple abstraction of a command line prompt for issuing game commands
*/
type CommandLinePrompt struct {
	Game   *game.Game
	Prompt *prompt.Prompt
}

/*
executor - Function for executing the logic of the command line prompt
*/
func executor(in string) {
	if in == "exit" {
		fmt.Println("Exiting simulation...")
		os.Exit(0)
	}
}

/*
completer - Function for providing auto-complete for the commands
*/
func completer(in prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{}
}

/*
NewCLI - Constructor for the CommandLinePrompt structure. Requires a pointer to the game object be passed as a constructor
*/
func NewCLI(game *game.Game) *CommandLinePrompt {
	return &CommandLinePrompt{
		Game:   game,
		Prompt: prompt.New(executor, completer, prompt.OptionPrefix(">>> ")),
	}
}

func (cmd *CommandLinePrompt) Start() {
	cmd.Prompt.Run()
}
