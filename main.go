package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"github.com/dispatchlabs/disgo_commons/crypto"
	"fmt"
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

	address, err := crypto.GetWalletAddress()
	if err != nil {
		fmt.Println(err.Error())
		//raw = NewWalletAddress()
	}


	log.Info(address)
}
