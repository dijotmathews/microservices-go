package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "dijo",
		Price: 4.33,
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
