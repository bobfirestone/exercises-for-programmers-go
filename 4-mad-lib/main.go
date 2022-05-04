package main

import "fmt"

func main() {
	inputs := []string{"noun", "verb", "adjective", "adverb"}
	m := make(map[string]string)
	for _, input := range inputs {
		fmt.Printf("Enter a %s: ", input)
		var i string
		fmt.Scanln(&i)
		m[input] = i
	}

	fmt.Printf("Do you %s your %s %s %s? That's hilarious!\n", m["verb"], m["adjective"], m["noun"], m["adverb"])
}
