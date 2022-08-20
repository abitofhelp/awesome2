package main

import (
	"fmt"
	"github.com/abitofhelp/awesome2/config"
	proxyrunner "github.com/abitofhelp/awesome2/proxy/runner"
	logger "github.com/labstack/gommon/log"
)

func main() {
	// Load environment variables.  If a filename is not provided, we are in production mode,
	// so we use the environment variables that exist on the host.
	// If a file is provided, we are in development mode, so we will load the configuration file,
	// but any environment variables that exist in the host will not be superseded.
	if appcfg, err := config.NewAppConfig("config/dev.env"); err == nil {
		if err := proxyrunner.Run(appcfg); err != nil {
			emsg := fmt.Errorf("\nfailed to start the server: %w", err)
			logger.Error(emsg)
			panic(emsg)
		}
	} else {
		panic(fmt.Errorf("\nfailed to load the application configuration for development: %w", err))
	}
}
