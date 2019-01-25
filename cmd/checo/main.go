package main

import (
	"os"
	"fmt"
	"github.com/nesheep5/checo"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: Required AccountName.     Usage: ./checo account_name")
		return
	}
	account := os.Args[1]
	checo.Run(account)
}
