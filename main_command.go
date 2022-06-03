package main

import (
	"fmt"
	"ydocker/container"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var initCmd = cli.Command{
	Name:  "init",
	Usage: "Init container process run user's process in container. Do not call it outside",
	Action: func(context *cli.Context) error {
		log.Infof("init come on")
		cmd := context.Args().Get(0)
		err := container.RunContainerInitProcess(cmd, nil)
		return err
	},
}

var runCmd = cli.Command{
	Name:  "run",
	Usage: "run container , ie: mydocker run -ti [image] [command]",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "ti",
			Usage: "enable tty",
		},
	},

	Action: func(ctx *cli.Context) error {
		if len(ctx.Args()) < 1 {
			return fmt.Errorf("missing container cmd")
		}

		cmd := ctx.Args().Get(0)
		tty := ctx.Bool("ti")

		Run(tty, cmd)

		return nil
	},
}
