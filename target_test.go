package main

import "testing"

func TestNewTarget_Empty(t *testing.T) {
	target, err := NewTarget("")
	if target.String() != "" {
		t.Errorf("expected target.String() to be empty, is %q", target)
	}
	if err != nil {
		t.Errorf("unexpected error: %q", err)
	}
}

func TestNewTarget_NoHost(t *testing.T) {
	in := "/"
	target, err := NewTarget(in)
	if target.IsAbs() {
		t.Error("unexpected absolute target")
	}
	if target.String() != in {
		t.Errorf("expected target.String() == %q, got %q", in, target)
	}
	if err != nil {
		t.Errorf("unexpected error: %q", err)
	}
}

func TestNewTarget_ProperURL(t *testing.T) {
	in := "http://www.example.com/jobs/33"
	target, err := NewTarget(in)
	if target.String() == "" {
		t.Error("expected target.URL.String() to not be empty")
	}
	if target.String() != in {
		t.Errorf("expected target.String() == %q, got %q", in, target)
	}
	if err != nil {
		t.Errorf("unexpected error: %q", err)
	}
}
