package cli

import "github.com/urfave/cli/v2"

// userCommand команды пользователей
func userCommand(app *App) *cli.Command {
	return &cli.Command{
		Name:    "user",
		Aliases: []string{"u"},
		Usage:   "User-related commands",
		Subcommands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "Add a new user",
				Flags:   userFlags(),
				Action:  app.AddUser,
			},
		},
	}
}
