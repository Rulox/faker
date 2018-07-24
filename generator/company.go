package generator

// CompanyData is the struct with the company YAML data
type CompanyData struct {
	Company struct {
		Suffix     []string   `yaml:"suffix"`
		Buzzwords  [][]string `yaml:"buzzwords"`
		Industry   []string   `yaml:"industry"`
		Profession []string   `yaml:"profession"`
		Formatters []string   `yaml:"formatters"`
	}
}

type CompanyGenerator struct {
	data CompanyData
}