package main

import (
	"fmt"
)

const (
	spanish  = "Spanish"
	french   = "French"
	japanese = "Japanese"

	englishHelloPrefix  = "Hello"
	spanishHelloPrefix  = "Hola"
	frenchHelloPrefix   = "Bonjour"
	japaneseHelloPrefix = "Konichiwa"
)

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}
	helloPrefix := greetingPrefix(lang)
	return fmt.Sprintf("%s, %s", helloPrefix, name)
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case japanese:
		prefix = japaneseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func MyName(name string) string {
	return name
}

func AddNums(num1 int64, num2 float64) float64 {
	return float64(num1) + num2
}

func main() {
	fmt.Println(Hello("Hamidou", ""))
	fmt.Println(Hello("Alpha", "Spanish"))
	fmt.Println(Hello("Ghost", "French"))
	fmt.Println(Hello("Rama", "Japanese"))
}
