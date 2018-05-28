package generator

import (
	"fmt"
	"math/rand"
	"strconv"
)

const DefaultLocales = "./locales"
const DefaultYamlName = "faker.yml"
const DefaultDigit = "#"

func GetYamlPath(lc string) string {
	return fmt.Sprintf("%s/%s/%s", DefaultLocales, lc, DefaultYamlName)
}

// Helper function that returns all the appearances of "#" for a random digit 0-9
func FormatDigits(s string) string {
	var r string
	for _, c := range s {
		if string(c) == DefaultDigit {
			r += strconv.Itoa(rand.Intn(9))
		} else {
			r += string(c)
		}
	}
	return r
}

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
	Misc 	MiscGenerator
	// Address data generator
	Address AddressGenerator
}

// Set locale for all the generators
func (f *Faker) SetLocale(l string) error {
	err := f.Address.supplyWithLocale(l)
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


