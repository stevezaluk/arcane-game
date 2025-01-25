package prompt

import (
	"github.com/c-bata/go-prompt"
	"github.com/stevezaluk/arcane-game/game"
)

// ExecuteFunction - The primary logic for executing the command
type ExecuteFunction func(game *game.Game, args []string)

/*
InteractiveCommand - Represents a command in the interactive command prompt
*/
type InteractiveCommand struct {
	Command     string
	Description string

	MinimumArgs int

	Execute ExecuteFunction
}

/*
ToPromptSuggest - Convert the Command to a prompt suggestion provided by go-prompt
*/
func (cmd *InteractiveCommand) ToPromptSuggest() prompt.Suggest {
	return prompt.Suggest{Text: cmd.Command, Description: cmd.Description}
}
