package main

import "fmt"

const englishHelloPrefix = "Hello, "
const SpanishSelector = "Spanish"
const spanishHelloPrefix = "Hola, "

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}
	if language == SpanishSelector {
		return spanishHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
