package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
	LanguageName() string
	Greet(visitorName string) string
}

type Italian struct{}

func (visitor Italian) LanguageName() string {
	return "Italian"
}

func (visitor Italian) Greet(visitorName string) string {
	return fmt.Sprintf("Ciao %s!", visitorName)
}

type Portuguese struct{}

func (visitor Portuguese) LanguageName() string {
	return "Portuguese"
}

func (visitor Portuguese) Greet(visitorName string) string {
	return fmt.Sprintf("Olá %s!", visitorName)
}

func SayHello(visitorName string, visitor Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", visitor.LanguageName(), visitor.Greet(visitorName))
}
