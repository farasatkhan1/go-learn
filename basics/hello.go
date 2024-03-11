package main

import "fmt"


func hello(name string, language string) string {
	const englishHelloPrefix = "Hello, "
	const spanishHelloPrefix = "Hola, "
	const frenchHelloPrefix = "Bonjour, "

	const spanish = "Spanish"
	const french = "French"

	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
		case spanish:
			prefix = spanishHelloPrefix
		case french:
			prefix = frenchHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(hello("Chris", "English"))
}