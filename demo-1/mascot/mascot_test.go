package mascot_test

import (
	"testing"

	"github.com/kunalrukade/Gocourse/demo-1/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Go Gopher" {
		t.Fatal("wrong mascot :(")
	}

}
