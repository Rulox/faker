package main

import (
	"github.com/Rulox/faker"
	"fmt"
)

// This is an example in a main function of your Golang application. For more specific
// examples refer to the other files in this folder!
func main() {
	f := faker.NewFaker("es_ES")
	fmt.Println(fmt.Sprintf("Address example: %s", f.Address.Full()))
	fmt.Println(fmt.Sprintf("Person Example: %s", f.Person.FullName()))
}