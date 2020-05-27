package model

type Config struct {
	Domains []struct {
		Domain   string `yaml:"domain"`
		Carriers []struct {
			Carrier   string `yaml:"carrier"`
			Selectors []struct {
				Type      string   `yaml:"type"`
				Selector  string   `yaml:"selector"`
				Start     int      `yaml:"start,omitempty"`
				End       int      `yaml:"end,omitempty"`
				Query     string   `yaml:"query"`
				Predicate string   `yaml:"predicate"`
				Values    []string `yaml:"values,omitempty"`
			} `yaml:"selectors"`
		} `yaml:"carriers"`
	} `yaml:"domains"`
}
