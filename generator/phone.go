package generator

import (
	"gopkg.in/yaml.v2"
	"math/rand"
)

// PhoneData is the struct with the company YAML data
type PhoneData struct {
	Phone struct {
		Formatters  []string `yaml:"formatters"`
	}
}

// PhoneGenerator the generator for phone numbers
type PhoneGenerator struct {
	data PhoneData
	// Whether use random formatter for full person name all the time or not
	RandomFormatter bool
	// SelectedFormatter (only if `randomFormatter` is false)
	SelectedFormatter string
}

// SetLocale sets the locale for the generator
func (pg *PhoneGenerator) SetLocale(lc string) error {
	y, err := GetYaml(lc)
	if err != nil {
		return err
	}

	pd := &PhoneData{}
	err = yaml.Unmarshal(y, &pd)
	if err != nil {
		return err
	}
	pg.data = *pd
	return nil
}

// PhoneNumber returns a phone number
func (pg *PhoneGenerator) PhoneNumber() string {
	f := pg.SelectedFormatter
	if !pg.RandomFormatter {
		f = pg.data.Phone.Formatters[rand.Intn(len(pg.data.Phone.Formatters))]
	}
	return FormatDigits(f)
}
