package main

import (
	"fmt"
)

const greetPrefix = "Hi, "
const sentenceEnd = ".The world need your help"
const botName = "botcito"

func main() {
	fmt.Println(Greet("go"))
}

func Greet(name string) string {
	if name == "" {
		return greetPrefix + "my name is " + botName + ".Can I know your name?"
	}

	return greetPrefix + name + sentenceEnd
}
