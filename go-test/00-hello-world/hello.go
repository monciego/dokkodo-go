package main

import "fmt"

const (
	spanish  = "Spanish"
	french   = "French"
	filipino = "Filipino"

	englishGreeting  = "Hello, "
	spanishGreeting  = "Hola, "
	frenchGreeting   = "Bonjour, "
	filipinoGreeting = "Kumusta, "
)

func Hello(name, language string) string {
	if name == "" {
		return englishGreeting + "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishGreeting
	case french:
		prefix = frenchGreeting
	case filipino:
		prefix = filipinoGreeting
	default:
		prefix = englishGreeting
	}

	return
}

func main() {
	fmt.Println(Hello("Jericho", "English"))
}
