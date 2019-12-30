package main

import "testing"

var botGreetText string = "Hi, my name is botcito.Can I know your name?: \n"

func assertMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf(" got %q want %q", got, want)
	}
}

func TestGreet(t *testing.T) {

	t.Run("Greet user when the name is provided", func(t *testing.T) {
		got := Greet("go")
		want := "Hi, go.The world need your help"

		assertMessage(t, got, want)
	})

	t.Run("Say bot name and ask for username when name is not provided", func(t *testing.T) {
		got := Greet("")
		want := botGreetText

		assertMessage(t, got, want)
	})

}

func TestBotGreet(t *testing.T) {
	t.Run("The bot shuld ask for the user name", func(t *testing.T) {
		got := BotGreet()
		want := botGreetText

		assertMessage(t, got, want)
	})
}
