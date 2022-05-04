package main

import (
	"bufio"
	"fmt"
	"os"
)

type quote struct {
	quote  string
	author string
}

func main() {
	q := quote{}
	fmt.Print("What is the quote? ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		q.quote = scanner.Text()
	}

	fmt.Print("Who said it? ")
	if scanner.Scan() {
		q.author = scanner.Text()
	}
	fmt.Printf("%s said, \"%s\"\n", q.author, q.quote)
}
