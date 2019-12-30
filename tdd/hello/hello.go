package main

import (
	"bufio"
	"fmt"
	"os"
)

const greetPrefix = "Hi, "
const sentenceEnd = ".The world need your help"
const botName = "botcito"

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println(BotGreet())
	name, _ := inputReader.ReadString('\n')
	fmt.Println(Greet(name))
}

func BotGreet() string {
	return greetPrefix + "my name is " + botName + ".Can I know your name?: \n"
}

func Greet(name string) string {
	greet := BotGreet()
	if name != "" {
		greet = greetPrefix + name + sentenceEnd
	}

	return greet
}
