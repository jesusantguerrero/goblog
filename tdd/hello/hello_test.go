package main

import "testing"

func TestGreet(t *testing.T) {

	assertMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf(" got %q want %q", got, want)
		}
	}

	t.Run("Greet user when the name is provided", func(t *testing.T) {
		got := Greet("go")
		want := "Hi, go.The world need your help"

		assertMessage(t, got, want)
	})

	t.Run("Say bot name and ask for username when name is not provided", func(t *testing.T) {
		got := Greet("")
		want := "Hi, my name is botcito.Can I know your name?"

		assertMessage(t, got, want)
	})
}
