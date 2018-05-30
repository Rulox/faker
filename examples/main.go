package main

import (
	"github.com/rulox/faker"
	"fmt"
)

// This is an example in a main function of your Golang application. For more specific
// examples refer to the other files in this folder!
func main() {
	f := faker.NewFaker("es_ES")
	fmt.Println(fmt.Sprintf("Address example: %s", f.Address.Full()))
}