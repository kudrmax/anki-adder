package config

import (
	"fmt"
	"io"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Deck      string   `yaml:"deck"`
	NoteModel string   `yaml:"noteModel"`
	Fields    []string `yaml:"fields"`
	DBFile    string   `yaml:"db_file"`
}

// Parse читает YAML-конфиг из io.Reader
func Parse(r io.Reader) (Config, error) {
	var cfg Config
	d := yaml.NewDecoder(r)
	if err := d.Decode(&cfg); err != nil {
		return Config{}, fmt.Errorf("error parsing config: %w", err)
	}

	return cfg, nil
}
