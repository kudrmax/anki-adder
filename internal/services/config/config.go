package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Deck      string   `yaml:"deck"`
	NoteModel string   `yaml:"noteModel"`
	Fields    []string `yaml:"fields"`
}

func Parse(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, err

	}

	return cfg, nil
}
