package cli

import (
	"github.com/urfave/cli/v2"
)

func (app *App) commands() []*cli.Command {
	return []*cli.Command{
		taskCommand(app),
		userCommand(app),
	}
}
