package main

import "fmt"

// struct Person represents a person with a name and age.
type Person struct {
	Name string
}

// Hello returns a greeting for the named person.
func Hello(p Person) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Selamat pagi!", p.Name)
	return message
}
