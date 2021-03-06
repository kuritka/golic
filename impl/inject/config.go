package inject

type Config struct {
	Golic struct{
		Licenses map[string]string `yaml:"licenses"`
		Rules    map[string]struct {
			Prefix string `yaml:"prefix"`
			Suffix string `yaml:"suffix"`
			Under []string `yaml:"under"`
		} `yaml:"rules"`
	} `yaml:"golic"`
}


func (c *Config) IsWrapped(key string) bool {
	return c.Golic.Rules[key].Suffix != ""
}

func (c *Config) LicenseStartsAfterHeader(key string) bool {
	return len(c.Golic.Rules[key].Under) > 0
}