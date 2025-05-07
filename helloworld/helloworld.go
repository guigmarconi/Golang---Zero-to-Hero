package main

import (
	"fmt"
)

const (
	portuguese = "Portuguese"
	french     = "French"
	spanish    = "Spanish"

	englishHelloPrefix = "Hello, "
	portuguesePrefix   = "Ola, "
	frenchPrefix       = "Bonjour, "
	spaninshPrefix     = "Hola, "
)

func Hello(name, language string) string {

	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case portuguese:
		prefix = portuguesePrefix
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spaninshPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Gustave", portuguese))
}
