package main

import "testing"

func TestTarget(t *testing.T) {
	var in string
	var err error

	in = ">|O.^.o|<"
	if _, err = NewTarget(in); err == nil {
		t.Error("expected an error.")
	}

	// breaks compile on bad type change.
	var act Target

	in = "http://www.example.com/"
	if act, err = NewTarget(in); err != nil {
		t.Error(err)
	}

	if act.String() != in {
		t.Errorf(
			"expected targ.String %q to match input %q.",
			act.String(), in)
	}
}
