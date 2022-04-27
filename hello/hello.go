package hello

const spanish = "Spanish"
const french = "French"
const frenchHelloPrefix = "Bonjour, "
const spanishHelloPrefix = "Hola, "
const englishHelloPrefix = "Hello, "

// Hello public function
func Hello(name string, language string) string {
	if name == "" {
		name = "worl"
	}

	return helloPrefix(language) + name

}

// helloPrefix for separating the component of helloPrefix
// this is private function
func helloPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return

}
