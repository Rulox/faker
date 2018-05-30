package faker

import (
	"fmt"
	"time"

	"math/rand"
	"github.com/rulox/faker/generator"
)

const defaultLocale = "en_US"

// Base struct for all generators.
type Faker struct {
	// If the generator should return always unique values. Note that this will have some limitations.
	// If you want more data than the available unique data, the Generator won't fail but it'll print a warning to
	// let you know it was impossible to generate unique data for you.
	Unique bool
	// Map to store the values in case unique is set to true
	used map[string]interface{}
	// Locale set. Depending on the locale, the generators will return the data in the selected language
	locale string

	// Misc data generator
	Misc 	generator.MiscGenerator
	// Address data generator
	Address generator.AddressGenerator
}

// Return a new Faker to start working with. It is necessary to use this 'constructor' in order
// to initialize some variables and to run the default locale `en_US`
func NewFaker(l string) (f Faker) {
	rand.Seed(time.Now().Unix())
	f.locale = l
	if len(l) == 0 {
		f.locale = defaultLocale
	}
	f.SetLocale(f.locale)
	return f
}

// Set locale for all the generators
func (f *Faker) SetLocale(l string) error {
	f.locale = l
	err := f.Address.SetLocale(l)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Clear the unique values storage so values can be repeated again
func (f *Faker) Clear() {
	f.used = make(map[string]interface{})
}


