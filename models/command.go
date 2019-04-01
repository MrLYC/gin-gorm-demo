package models

import (
	"context"
	"flag"

	"github.com/google/subcommands"
)

// Command :
type Command struct{}

// Name :
func (*Command) Name() string {
	return "migrate"
}

// Synopsis :
func (*Command) Synopsis() string {
	return "Auto migrate models"
}

// Usage :
func (*Command) Usage() string {
	return `migrate
  Auto migrate models.
`
}

// SetFlags :
func (*Command) SetFlags(f *flag.FlagSet) {
}

// Execute :
func (p *Command) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	err := OpenDB()
	if err != nil {
		panic(err)
	}
	defer CloseDB()

	// Migrate the schema
	DB.AutoMigrate(&Pair{})

	return subcommands.ExitSuccess
}
