package demo

import (
	"context"
	"flag"
	"fmt"

	"github.com/mrlyc/gin-gorm-demo/models"

	"github.com/gin-gonic/gin"
	"github.com/google/subcommands"
	"github.com/mrlyc/gin-gorm-demo/config"
	"github.com/mrlyc/gin-gorm-demo/logging"
)

// Command :
type Command struct{}

// Name :
func (*Command) Name() string {
	return "server"
}

// Synopsis :
func (*Command) Synopsis() string {
	return "Run server"
}

// Usage :
func (*Command) Usage() string {
	return `server
  Run server.
`
}

// SetFlags :
func (*Command) SetFlags(f *flag.FlagSet) {

}

// Execute :
func (p *Command) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	logger := logging.GetLogger()
	conf := config.Configuration.HTTP
	address := fmt.Sprintf("%v:%v", conf.Host, conf.Port)

	engine := gin.Default()
	route(engine)

	err := models.OpenDB()
	if err != nil {
		panic(err)
	}

	models.DB.SetLogger(logger)

	if !config.Configuration.Debug {
		gin.SetMode(gin.ReleaseMode)
		models.DB.Debug()
	}

	logger.Infof("gin-gorm-demo listening on %v", address)
	err = engine.Run(address)
	if err != nil {
		panic(err)
	}
	return subcommands.ExitSuccess
}
