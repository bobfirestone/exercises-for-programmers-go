package main

import "fmt"

func main() {
	var name string
	var greeting string
	fmt.Print("What is your name? ")
	fmt.Scanf("%s", &name)
	switch name {
	case "Alice":
		greeting = "nice to see you again."
	case "Bob":
		greeting = "nice to meet you!"
	default:
		greeting = "why are you here?"

	}
	fmt.Printf("Hello %s, %s\n", name, greeting)
	return
}
