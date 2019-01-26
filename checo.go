package checo

import (
	"fmt"
	"net/http"
)

var Checkers = []Checker{
	Checker{Name: "Twitter", Url: "https://twitter.com/"},
	Checker{Name: "Facebook", Url: "https://www.facebook.com/"},
	Checker{Name: "Instagram", Url: "https://www.instagram.com/"},
}

type Checker struct {
	Name string
	Url  string
}

func (c Checker) exists(account string) (bool, error) {
	resp, err := http.Get(c.Url + account)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	return resp.StatusCode == 200, nil
}

func Run(account string) error {
	if account == "" {
		fmt.Errorf("account is required.")
	}

	fmt.Printf("Search Account: %v \n\n", account)

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

		fmt.Printf("%v : %v \n", c.Name, msg)
	}
	return nil
}
