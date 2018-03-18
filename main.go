package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Version = Version
	app.Author = "morphtoken"
	app.Name = "MorphToken CLI"
	app.Email = "contact@morphtoken.com"
    app.Usage = "Exchange coins instantly from your terminal"

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	app.Run(os.Args)
}
