package models

import (
	"context"
	"flag"
	"github.com/google/subcommands"
	"github.com/jinzhu/gorm"
	"github.com/mrlyc/gin-gorm-demo/config"
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
	db, err := gorm.Open(config.Configuration.Database.Type, config.Configuration.Database.Connection)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Pairs{})

	return subcommands.ExitSuccess
}
