package hello

const englishHelloPrefix = "Hello, "
const Spanish = "Spanish"
const spanishHelloPrefix = "Hola, "
const French = "French"
const frenchHelloPrefix = "Bonjour, "
const Dutch = "Dutch"
const dutchHelloPrefix = "Hallo, "

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case French:
		prefix = frenchHelloPrefix
	case Spanish:
		prefix = spanishHelloPrefix
	case Dutch:
		prefix = dutchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix
}
