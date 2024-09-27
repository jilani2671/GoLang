package main

import "fmt"

const greetings = "Hello,"

func Greetings(name string) string {

	if name == "" {
		name = "World"
	}

	return greetings + name
}

func main() {
	fmt.Println(Greetings(""))
}

// Welcome to GoLang Mr.Abrar
