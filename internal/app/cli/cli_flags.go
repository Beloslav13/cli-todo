package cli

import "github.com/urfave/cli/v2"

func taskFlags() []cli.Flag {
	return []cli.Flag{
		&cli.Int64Flag{Name: "user_id", Usage: "User ID for the task", Required: true},
		&cli.StringFlag{Name: "name", Usage: "Task name", Required: true},
		&cli.StringFlag{Name: "status", Usage: "Task status (new, in_progress, completed)", Required: true},
	}
}

func listTaskFlags() []cli.Flag {
	return []cli.Flag{
		&cli.Int64Flag{Name: "user_id", Usage: "List tasks by user ID", Required: true},
		&cli.StringFlag{Name: "status", Usage: "Filter tasks by status (new, in_progress, completed)", Required: false},
		&cli.StringFlag{Name: "sort", Usage: "Sort tasks by id or created_at", Value: "created_at"},
		&cli.StringFlag{Name: "order", Usage: "Sort order (asc, desc)", Value: "desc"},
	}
}

func ListAllTasksFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{Name: "status", Usage: "Filter tasks by status (new, in_progress, completed)", Required: false},
	}
}

func userFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{Name: "username", Usage: "Username", Required: true},
	}
}
