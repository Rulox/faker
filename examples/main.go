package main

import (
	"github.com/Rulox/faker"
	"fmt"
)

// This is an example in a main function of your Golang application.
func main() {
	f := faker.NewFaker("es_ES")
	fmt.Println(fmt.Sprintf("Address example: %s", f.Address.Full()))
	fmt.Println(fmt.Sprintf("Person Example: %s", f.Person.FullName()))
	// Some phones
	fmt.Println(f.Phone.PhoneNumber())
}