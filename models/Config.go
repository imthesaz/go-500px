package models

type Config struct {
	SearchConfig SearchConfig `yaml:"search_config"`
}

type SearchConfig struct {
	SearchTerm string   `yaml:"search_term"`
	Sort       []string `yaml:"sort"`
	Count      int      `yaml:"count"`
}
