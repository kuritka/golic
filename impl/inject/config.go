package inject

type Config struct {
	Golic struct{
		Licenses map[string]string `yaml:"licenses"`
		Comments map[string]struct {
			Prefix string `yaml:"prefix"`
			Suffix string `yaml:"suffix"`
		} `yaml:"comments"`
	} `yaml:"golic"`
}
