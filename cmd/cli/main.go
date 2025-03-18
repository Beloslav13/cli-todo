package main

import (
	"fmt"
	"github.com/Beloslav13/cli-todo/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// init config
	cfg := config.MustLoad[config.CLIConfig]()
	fmt.Printf("%+v\n", cfg)

	// logger
	// init app
	// start app

}
