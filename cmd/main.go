package main

import (
	"os"

	"github.com/nesheep5/checo"
)

func main() {
	account := os.Args[1]
	checo.Run(account)
}
