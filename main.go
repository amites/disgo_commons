package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"github.com/dispatchlabs/disgo_commons/types"
)

func main() {
	// Setup log.
	formatter := &log.TextFormatter{
		FullTimestamp: true,
		ForceColors:   false,
	}
	log.SetFormatter(formatter)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	account, error := types.NewAccount()
	if error != nil {

	}

	log.Info(account)
}
