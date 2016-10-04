package main

import (
	"os"

	"github.com/ckeyer/monitor/server"
	"gopkg.in/urfave/cli.v2"
)

func main() {
	var addr string

	serveCmd := &cli.Command{
		Name:  "serve",
		Usage: "start a server for prometheus.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "addr",
				Aliases:     []string{"serverAddr"},
				Usage:       "[host]:port",
				EnvVars:     []string{"SERVER_ADDR"},
				Value:       ":1993",
				Destination: &addr,
			},
		},
		Action: func(ctx *cli.Context) error {
			server.Serve(addr)
			return nil
		},
	}
	clientCmd := &cli.Command{
		Name:    "client",
		Aliases: []string{"serverAddr"},
		Usage:   "start a client.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "addr",
				Usage:       "[host]:port",
				EnvVars:     []string{"SERVER_ADDR"},
				Value:       ":1993",
				Destination: &addr,
			},
		},
		Action: func(ctx *cli.Context) error {
			return nil
		},
	}

	app := &cli.App{
		Name:    "monitor",
		Version: "0.2",
		Commands: []*cli.Command{
			serveCmd,
			clientCmd,
		},
	}
	app.Run(os.Args)
}
