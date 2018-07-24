package generator

import (
	"gopkg.in/yaml.v2"
	"math/rand"
	"strings"
)
// PersonData is the struct for the person YAML data
type PersonData struct {
	Person struct {
		FirstName  []string `yaml:"first_name"`
		MiddleName []string `yaml:"middle_name"`
		LastName   []string `yaml:"last_name"`
		Age        []string `yaml:"age"`
		Formatters []string `yaml:"formatters"`
	}
}

// PersonGenerator is the Generator for People.
type PersonGenerator struct {
	data              PersonData
	// Whether use random formatter for full person name all the time or not
	RandomFormatter   bool
	// SelectedFormatter (only if `randomFormatter` is false)
	SelectedFormatter string
}

// Supply the formatter with the right data using a locale. By default en_US
func (pg *PersonGenerator) SetLocale(lc string) error {
	y, err := GetYaml(lc)
	if err != nil {
		return err
	}

	pd := &PersonData{}
	err = yaml.Unmarshal(y, &pd)
	if err != nil {
		return err
	}
	pg.data = *pd
	return nil
}

func (pg *PersonGenerator) formatFullName() string {
	f := pg.SelectedFormatter
	if !pg.RandomFormatter {
		f = pg.data.Person.Formatters[rand.Intn(len(pg.data.Person.Formatters))]
	}
	var r string
	elems := strings.Split(f, " ")
	for i, w := range elems {
		switch w {
		case "{{first_name}}":
			r += pg.FirstName()
		case "{{middle_name}}":
			r += pg.MiddleName()
		case "{{last_name}}":
			r += pg.LastName()
		}
		if i != len(elems) {
			r += " "
		}
	}
	return r
}

// Get a full name using a formatter
func (pg *PersonGenerator) FullName() string {
	return pg.formatFullName()
}
// Get person First Name
func (pg *PersonGenerator) FirstName() string {
	return pg.data.Person.FirstName[rand.Intn(len(pg.data.Person.FirstName))]
}
// Get person Last Name
func (pg *PersonGenerator) LastName() string {
	return pg.data.Person.LastName[rand.Intn(len(pg.data.Person.LastName))]
}
// Get person Middle Name
func (pg *PersonGenerator) MiddleName() string {
	return pg.data.Person.MiddleName[rand.Intn(len(pg.data.Person.MiddleName))]
}
// Get person age
func (pg *PersonGenerator) Age() string {
	return FormatDigits(pg.data.Person.Age[rand.Intn(len(pg.data.Person.Age))])
}