package main

import (
	"fmt"

	"github.com/Alevsk/go-app-template/pkg"
	"github.com/minio/cli"
)

// starts the server
var serverCmd = cli.Command{
	Name:    "server",
	Aliases: []string{"srv"},
	Usage:   "Start App server",
	Action:  StartServer,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "",
			Usage: "bind to a specific HOST, HOST can be an IP or hostname",
		},
		cli.IntFlag{
			Name:  "port",
			Value: 8080,
			Usage: "bind to specific HTTP port",
		},
	},
}

// StartServer starts the app server
func StartServer(ctx *cli.Context) error {

	pkg.LogInfo(ctx.String("host"))
	pkg.LogInfo(fmt.Sprintf("%d", ctx.Int("port")))

	return nil
}
