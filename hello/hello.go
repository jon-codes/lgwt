package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	englishHelloSuffix = "!"

	spanishHelloPrefix = "¡Hola, "
	spanishHelloSuffix = "!"

	frenchHelloPrefix = "Bonjour, "
	frenchHelloSuffix = "!"
)

func greetingPrefixSuffix(lang string) (string, string) {
	switch lang {
	case spanish:
		return spanishHelloPrefix, spanishHelloSuffix
	case french:
		return frenchHelloPrefix, frenchHelloSuffix
	default:
		return englishHelloPrefix, englishHelloSuffix
	}
}

func Hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}
	prefix, suffix := greetingPrefixSuffix(lang)
	return prefix + name + suffix
}

func main() {
	fmt.Println(Hello("Jon", "English"))
}
