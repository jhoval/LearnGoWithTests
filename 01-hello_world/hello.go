package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const german = "German"
const englishHelloPrefish = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const germanHelloPrefix = "Hallo, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	} 

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language{
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case german:
		prefix = germanHelloPrefix
	default:
		prefix = englishHelloPrefish
	}
	return
}


func main() {
	fmt.Println(Hello("world","English"))
}
