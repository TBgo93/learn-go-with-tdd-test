package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefic = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return grettingPrefix(language) + name
}

func grettingPrefix(language string) (prefix string) {
	switch language {
	case "French":
		prefix = frenchHelloPrefic
	case "Spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
