package main

const (
	french  = "French"
	spanish = "Spanish"
	english = "English"

	spanishHelloPrefix = "Hola, "

	englishHelloPrefix = "Hello, "
	frenchHelloPrefix  = "Allo, "
)

func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case french:
		prefix = frenchHelloPrefix

	case spanish:
		prefix = spanishHelloPrefix

	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
}
