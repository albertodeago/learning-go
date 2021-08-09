package main

import "fmt"

const spanish = "Spanish"
const french = "French"

const englishHello = "Hello, "
const spanishHello = "Hola, "
const frenchHello = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHello
	case french:
		prefix = frenchHello
	default:
		prefix = englishHello
	}

	return
}

func main() {
	fmt.Println(Hello("Man", ""))
}
