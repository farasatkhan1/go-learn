package main

import "fmt"


func hello(name string) string {
	const englishHelloPrefix = "Hello, "
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(hello("Chris"))
}