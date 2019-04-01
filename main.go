package main

import (
	"context"
	"flag"
	"os"

	"github.com/mrlyc/gin-gorm-demo/models"

	"github.com/google/subcommands"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mrlyc/gin-gorm-demo/config"
	"github.com/mrlyc/gin-gorm-demo/gin-gorm-demo"
)

type initialHandler func() bool

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&config.VersionCommand{}, "")
	subcommands.Register(&config.ConfInfoCommand{}, "")
	subcommands.Register(&demo.Command{}, "")
	subcommands.Register(&models.Command{}, "")

	flag.StringVar(
		&(config.Configuration.ConfigurationPath),
		"c", config.Configuration.ConfigurationPath,
		"Configuration file",
	)

	flag.Parse()

	initialHandlers := []initialHandler{
		initRandomSeed,
		initConfiguration,
	}

	for _, handler := range initialHandlers {
		if !handler() {
			os.Exit(255)
		}
	}

	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
