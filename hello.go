package main

import "fmt"

const greetingName = "Hello,"

func Greetings(name string) string {
	return greetingName + name
}

// Hello,Abrar

func main() {
	fmt.Println(Greetings("Abrar"))
}
