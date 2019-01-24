package checo

import (
	"fmt"
	"net/http"
)

var Checkers = []Checker{
	Checker{name: "Twitter", url: "https://twitter.com/"},
}

type Checker struct {
	name string
	url  string
}

func (c Checker) exists(account string) (bool, error) {
	resp, err := http.Get(c.url + account)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	return resp.StatusCode == 200, nil
}

func Run(account string) error {
	for _, c := range Checkers {
		exists, err := c.exists(account)
		if err != nil {
			return err
		}
		var msg string
		if exists {
			msg = "Oops! Alrady Exists..."
		} else {
			msg = "OK! No Exists!"
		}

		fmt.Printf("%v : %v", c.name, msg)
	}
	return nil
}
