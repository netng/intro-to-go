package hello

import (
	"fmt"
)

const spanish = "Spanish"
const french = "French"
const frenchHelloPrefix = "Bonjour, "
const spanishHelloPrefix = "Hola, "
const englishHelloPrefix = "Hello, "

// Hello public function
func Hello(name string, language string) string {
	if name == "" {
		name = "world"
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

// main, this is just a main function
func main() {
	fmt.Println(Hello("", ""))
	fmt.Println(Hello("jono", "French"))
	fmt.Println(Hello("jono", "Spanish"))
	fmt.Println(Hello("jono", "English"))
	fmt.Println(Hello("jono", ""))
}
