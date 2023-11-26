package main

import (
	"fmt"
)

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

type German struct{}

func (g German) LanguageName() string {
	return "German"
}

func (g German) Greet(visitorName string) string {
	return fmt.Sprintf("Hallo %s!", visitorName)
}

type Italian struct{}

func (g Italian) LanguageName() string {
	return "Italian"
}

func (g Italian) Greet(visitorName string) string {
	return fmt.Sprintf("Ciao %s!", visitorName)
}

type Portuguese struct{}

func (g Portuguese) LanguageName() string {
	return "Portuguese"
}

func (g Portuguese) Greet(visitorName string) string {
	return fmt.Sprintf("Olá %s!", visitorName)
}

func SayHello(visitorName string, greeter Greeter) string {
	language := greeter.LanguageName()
	greeting := greeter.Greet(visitorName)
	return fmt.Sprintf("I can speak %s: %s", language, greeting)
}

func main() {
	german := German{}
	greeting := SayHello("Dietrich", german)
	fmt.Println(greeting)
}
