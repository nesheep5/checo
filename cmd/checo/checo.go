package main

import (
	"fmt"
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

	app.Commands = []cli.Command{
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "show checking SNS list",
			Action: func(c *cli.Context) error {
				fmt.Println("checking SNS list:")
				for _, ch := range checo.Checkers {
					fmt.Printf("  %v", ch.Name)
				}
				return nil
			},
		},
		{
			Name:    "check",
			Aliases: []string{"c"},
			Usage:   "checking SNS account",
			Action: func(c *cli.Context) error {
				account := c.Args().Get(0)
				return checo.Run(account)
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
