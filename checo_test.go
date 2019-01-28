package checo

import "testing"

type TestData struct {
	SNS     string
	Account string
	Actual  bool
}

func TestExists(t *testing.T) {
	tds := []TestData{
		TestData{"Twitter", "nesheep5", true},
		TestData{"Facebook", "shogo.watanabe.583", true},
		TestData{"Instagram", "shogo.mizuno", true},
	}
	for _, td := range tds {
		t.Run(td.SNS, func(t *testing.T) {
			b, err := CheckerMap[td.SNS].Exists(td.Account)
			if err != nil {
				t.Errorf("Occerd Error. err: %v", err)
			}
			if b != td.Actual {
				t.Errorf("'Exists' got %v, want %v", b, td.Actual)
			}
		})
	}
}
