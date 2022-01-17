package models

type Config struct {
	SearchConfig SearchConfig `yaml:"search_config"`
}

type SearchConfig struct {
	SearchTerm string   `yaml:"search_term"`
	Sort       []string `yaml:"sort"`
	Count      int      `yaml:"count"`
	PhotoIndex int      `yaml:"photo_index"`
	PrevSort   string   `yaml:"prev_sort"`
}
