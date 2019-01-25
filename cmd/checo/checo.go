package main

import (
	"log"
	"os"

	"github.com/nesheep5/checo"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "checo"
	app.Usage = "Function to check if an SNS account exists."
	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) error {
		if c.NArg() != 1 {
			cli.ShowAppHelpAndExit(c, 0)
		}

		account := c.Args().Get(0)
		return checo.Run(account)
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
