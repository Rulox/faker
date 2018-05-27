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
	// Selected locale
	locale 				string
	// Data loaded from YAML file for this locale
	data 				AddressData
	// Wether use random formatter for full address all the time or not
	randomFormatter 	bool
	// Selected formatter (only if `randomFormatter` is false)
	selectedFormatter	string
}
// Supply the formatter with the right data using a locale. By default en_US
func (ag *AddressGenerator) supplyWithLocale(lc string) error {
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
	// TODO work with commas in formatters
	for i, w := range elems {
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
			r += " "
		}
	}
	return r
}
// Add a custom formatter following the right structure. It can return an error on failure
//func (ag *AddressGenerator) AddFormatter(f string) error {}

func (ag *AddressGenerator) Full() string {
	return ag.formatFullAddress()
}
func (ag *AddressGenerator) City() string {
	return ag.data.Address.City[rand.Intn(len(ag.data.Address.City))]
}
func (ag *AddressGenerator) Country() string {
	return ag.data.Address.Country[rand.Intn(len(ag.data.Address.Country))]
}
func (ag *AddressGenerator) CountryCode() string {
	return ag.data.Address.CountryCode
}
//func (ag *AddressGenerator) Number() string {}
func (ag *AddressGenerator) StreetSuffix() string {
	return ag.data.Address.StreetSuffix[rand.Intn(len(ag.data.Address.StreetSuffix))]
}
func (ag *AddressGenerator) Street() string {
	return ag.data.Address.StreetName[rand.Intn(len(ag.data.Address.StreetName))]
}
func (ag *AddressGenerator) Province() string {
	return ag.data.Address.Province[rand.Intn(len(ag.data.Address.Province))]
}
func (ag *AddressGenerator) State() string {
	return ag.data.Address.State[rand.Intn(len(ag.data.Address.State))]
}
func (ag *AddressGenerator) Secondary() string {
	return FormatDigits(ag.data.Address.Secondary[rand.Intn(len(ag.data.Address.Secondary))])
}
func (ag *AddressGenerator) Number() string {
	return FormatDigits(ag.data.Address.Number[rand.Intn(len(ag.data.Address.Number))])
}
func (ag *AddressGenerator) ZipCode() string {
	return FormatDigits(ag.data.Address.ZipCode)
}
