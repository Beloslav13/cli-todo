package cli

import "github.com/urfave/cli/v2"

// taskCommand команды для tasks
func taskCommand(app *App) *cli.Command {
	return &cli.Command{
		Name:    "task",
		Aliases: []string{"t"},
		Usage:   "Task-related commands",
		Subcommands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "Add a task to the list",
				Flags:   taskFlags(),
				Action:  app.AddTask,
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "List tasks by user",
				Flags:   listTaskFlags(),
				Action:  app.ListTasksByUser,
			},
			{
				Name:    "alist",
				Aliases: []string{"al"},
				Usage:   "All list tasks",
				Flags:   ListAllTasksFlags(),
				Action:  app.ListAllTasks,
			},
		},
	}
}
