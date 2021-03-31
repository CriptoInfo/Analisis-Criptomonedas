package test

import (
	"testing"

	gecko "github.com/superoo7/go-gecko/v3"
)

func TestSimplePrice(t *testing.T) {
	cg := gecko.NewClient(nil)

	ids := []string{"bitcoin", "ethereum"}
	vc := []string{"usd", "eur"}
	_, err := cg.SimplePrice(ids, vc)
	if err != nil {
		t.Errorf("Invalid operation")
	}
}
