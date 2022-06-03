package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const usageMsg = `ydocker is a simple container toy level runtime implementation, which aims to learn
                       how docker works. enjoy to use ydocker`

func main() {
	app := cli.NewApp()
	app.Name = "ydocker"
	app.Usage = usageMsg

	app.Commands = []cli.Command{
		initCmd,
		runCmd,
	}

	app.Before = func(ctx *cli.Context) error {
		//set log formatter to JSON Formatter
		log.SetFormatter(&log.JSONFormatter{})
		log.SetOutput(os.Stdout)

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
