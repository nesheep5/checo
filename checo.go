package checo

import (
	"net/http"
)

var CheckerMap = map[string]Checker{
	"Twitter":   Checker{Name: "Twitter", Url: "https://twitter.com/"},
	"Facebook":  Checker{Name: "Facebook", Url: "https://www.facebook.com/"},
	"Instagram": Checker{Name: "Instagram", Url: "https://www.instagram.com/"},
}

type Checker struct {
	Name string
	Url  string
}

func (c Checker) Exists(account string) (bool, error) {
	resp, err := http.Get(c.Url + account)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	return resp.StatusCode == 200, nil
}
