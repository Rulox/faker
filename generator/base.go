package generator

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
	// Default MiscGenerator
	Misc MiscGenerator
}

func (f *Faker) New(unique bool) Faker {
	f.Misc = MiscGenerator{}
	return *f
}

// Clear the unique values storage so values can be repeated again
func (f *Faker) Clear() {
	f.used = make(map[string]interface{})
}

