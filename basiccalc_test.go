package basiccalc_test

import (
	"testing"
	. "github.com/Stanlyzoolo/basiccalc"
)

var input string = "2+1"

func TestEval(t *testing.T) {
	want := 3

	got, err := Eval(input)
	if got != want {
		t.Error("Something went wrong", err)
	}
}