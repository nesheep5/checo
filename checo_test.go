package checo

import "testing"

func TestExistsTwitter(t *testing.T) {
	b, err := CheckerMap["Twitter"].Exists("nesheep5")
	if err != nil {
		t.Errorf("Occerd Error. err: %v", err)
	}
	actual := true
	if b != actual {
		t.Errorf("'Exists' got %v, want %v", b, actual)
	}
}

func TestExistsFacebook(t *testing.T) {
	b, err := CheckerMap["Facebook"].Exists("shogo.watanabe.583")
	if err != nil {
		t.Errorf("Occerd Error. err: %v", err)
	}
	actual := true
	if b != actual {
		t.Errorf("'Exists' got %v, want %v", b, actual)
	}
}
func TestExistsInstagram(t *testing.T) {
	b, err := CheckerMap["Instagram"].Exists("shogo.mizuno")
	if err != nil {
		t.Errorf("Occerd Error. err: %v", err)
	}
	actual := true
	if b != actual {
		t.Errorf("'Exists' got %v, want %v", b, actual)
	}
}