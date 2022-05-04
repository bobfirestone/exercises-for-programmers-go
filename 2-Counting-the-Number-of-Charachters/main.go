package main

import "fmt"

func main() {
	fmt.Print("What is the input string? ")
	var input string
	fmt.Scanf("%s", &input)
	if len(input) > 0 {
		fmt.Printf("%s has %d characters\n", input, len(input))
	} else {
		fmt.Println("The input string is empty. Goodbye.")
	}

}
