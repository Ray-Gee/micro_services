package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name: "Ryu",
		Price: 1.9,
		SKU: "abs-dss-gee",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
