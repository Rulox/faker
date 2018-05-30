package generator

import (
	"path/filepath"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"math/rand"
	"strings"
)
type AddressData struct {
	Address struct {
		City           []string `yaml:"city"`
		Country        []string `yaml:"country"`
		CountryCode    string   `yaml:"country_code"`
		CountryDefault string   `yaml:"country_default"`
		Number         []string `yaml:"number"`
		StreetSuffix   []string `yaml:"street_suffix"`
		StreetPrefix   []string `yaml:"street_prefix"`
		StreetName 	   []string `yaml:"street_name"`
		Secondary	   []string `yaml:"secondary_address"`
		ZipCode        string   `yaml:"zip"`
		Province       []string `yaml:"province"`
		State          []string `yaml:"state"`
		Formatters     []string `yaml:"formatters"`
	}
}
// Generator for Addresses. It can return a fully address within the supported formats or just
type AddressGenerator struct {
	// Data loaded from YAML file for this locale
	data 				AddressData
	// Wether use random formatter for full address all the time or not
	randomFormatter 	bool
	// Selected formatter (only if `randomFormatter` is false)
	selectedFormatter	string
}

// Supply the formatter with the right data using a locale. By default en_US
func (ag *AddressGenerator) SetLocale(lc string) error {
	// FIX use a relative path
	f, err := filepath.Abs(GetYamlPath(lc))
	if err != nil {
		return err
	}

	y, err := ioutil.ReadFile(f)
	if err != nil {
		return err
	}

	ad := &AddressData{}
	err = yaml.Unmarshal(y, &ad)
	if err != nil {
		return err
	}
	ag.data = *ad
	return nil
}

// Create a full address within the registered formats
func (ag *AddressGenerator) formatFullAddress() string {
	f := ag.selectedFormatter
	if !ag.randomFormatter {
		f = ag.data.Address.Formatters[rand.Intn(len(ag.data.Address.Formatters))]
	}
	var r string
	elems := strings.Split(f, " ")

	for i, w := range elems {
		// Be sure there is no special character around the word. If there are, save them
		// to use them later
		var suffix string
		var prefix string
		if !strings.HasSuffix(w, "}}") {
			suffix = w[len(w)-1:]
			w = strings.TrimSuffix(w, suffix)
		}
		if !strings.HasPrefix(w, "{{") {
			prefix = w[0:1]
			w = strings.TrimPrefix(w, prefix)
		}
		switch w {
		case "{{street_suffix}}":
			r += ag.StreetSuffix()
		case "{{street_name}}":
			r += ag.Street()
		case "{{number}}":
			r += ag.Number()
		case "{{secondary_address}}":
			r += ag.Secondary()
		case "{{zip_code}}":
			r += ag.ZipCode()
		case "{{province}}":
			r += ag.Province()
		case "{{city}}":
			r += ag.City()
		case "{{state}}":
			r += ag.State()
		default:
			r += w
		}

		if i != len(elems) {
			// Add pre/su-ffixes to our word
			if len(suffix) != 0 {
				r += suffix
			}
			if len(prefix) != 0 {
				r = prefix + r
			}
			r += " "
		}
	}
	return r
}
// Add a custom formatter following the right structure. It can return an error on failure
//func (ag *AddressGenerator) AddFormatter(f string) error {}

// Get Full address using a formatter
func (ag *AddressGenerator) Full() string {
	return ag.formatFullAddress()
}
// Get a random City name
func (ag *AddressGenerator) City() string {
	return ag.data.Address.City[rand.Intn(len(ag.data.Address.City))]
}
// Get a random Country name
func (ag *AddressGenerator) Country() string {
	return ag.data.Address.Country[rand.Intn(len(ag.data.Address.Country))]
}
// Get a random Country code
func (ag *AddressGenerator) CountryCode() string {
	return ag.data.Address.CountryCode
}
// Get a random street suffix
func (ag *AddressGenerator) StreetSuffix() string {
	return ag.data.Address.StreetSuffix[rand.Intn(len(ag.data.Address.StreetSuffix))]
}
// Get a random street prefix
func (ag *AddressGenerator) StreetPrefix() string {
	return ag.data.Address.StreetPrefix[rand.Intn(len(ag.data.Address.StreetPrefix))]
}
// Get a random Street name
func (ag *AddressGenerator) Street() string {
	return ag.data.Address.StreetName[rand.Intn(len(ag.data.Address.StreetName))]
}
func (ag *AddressGenerator) Province() string {
	return ag.data.Address.Province[rand.Intn(len(ag.data.Address.Province))]
}
// Get a random State name
func (ag *AddressGenerator) State() string {
	return ag.data.Address.State[rand.Intn(len(ag.data.Address.State))]
}
// Get a random secondary address field
func (ag *AddressGenerator) Secondary() string {
	return FormatDigits(ag.data.Address.Secondary[rand.Intn(len(ag.data.Address.Secondary))])
}
// Get a random number using the formatter
func (ag *AddressGenerator) Number() string {
	return FormatDigits(ag.data.Address.Number[rand.Intn(len(ag.data.Address.Number))])
}
// Get a random zip code using the digits formatter
func (ag *AddressGenerator) ZipCode() string {
	return FormatDigits(ag.data.Address.ZipCode)
}
