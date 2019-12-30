package main

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("go")
	want := "Hello go"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
