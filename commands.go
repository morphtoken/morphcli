package main

import (
	"fmt"
	"os"

	"github.com/morphtoken/morphcli/command"
	"github.com/urfave/cli"
)

var GlobalFlags = []cli.Flag{}


var Commands = []cli.Command{
	{
		Name:   "rates",
		Usage:  "Get all instant rates",
		Action: command.CmdRates,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "exchange",
		Usage:  "Exchange one coin for another",
		Action: command.CmdExchange,
		Flags:  []cli.Flag{
            cli.StringFlag{
                Name: "input",
                Usage: "Input asset to convert from",
            },
            cli.StringFlag{
                Name: "output",
                Usage: "Output asset to convert to",
            },
            cli.StringFlag{
                Name: "address",
                Usage: "Address that will receive the output",
            },
            cli.StringFlag{
                Name: "refund",
                Usage: "Refund address to use in case it becomes necessary",
            },
        },
	},
    {
        Name:   "view",
        Usage:  "Fetch an existing trade",
        Action: command.CmdView,
        Flags: []cli.Flag{
            cli.StringFlag{
                Name: "id",
                Usage: "Morph trade to lookup, pass its id or the deposit address",
            },
        },
    },
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
