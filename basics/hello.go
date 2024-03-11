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

	if language == spanish {
		return spanishHelloPrefix + name
	}

	if language == french {
		return frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(hello("Chris", "English"))
}