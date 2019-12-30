package main

import (
	"fmt"
)

func main() {
	fmt.Println(Greet("go"))
}

func Greet(name string) string {
	return "Hello " + name
}
