package generator

import (
	"fmt"
	"strconv"
	"math/rand"
)
const defaultLocales = "./generator/locales"
const defaultYamlName = "faker.yml"
const defaultDigit = "#"

func GetYamlPath(lc string) string {
	return fmt.Sprintf("%s/%s/%s", defaultLocales, lc, defaultYamlName)
}

// Helper function that returns all the appearances of "#" for a random digit 0-9
func FormatDigits(s string) string {
	var r string
	for _, c := range s {
		if string(c) == defaultDigit {
			r += strconv.Itoa(rand.Intn(9))
		} else {
			r += string(c)
		}
	}
	return r
}
