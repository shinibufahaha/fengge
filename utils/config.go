package utils

import (
	"gopkg.in/yaml.v2"
	"sync"
	"os"
	"log"
)

type Config struct {
	// Server   ServerConfig   `yaml:"server"`
	// Database DatabaseConfig `yaml:"database"`
    Joern    JoernConfig    `yaml:"joernPath"`
	Path PathConfig `yaml:"filesPath"`
}

type JoernConfig struct {
    Joern       string  `yaml:"joern"`
    JoernParse  string  `yaml:"joernParse"`
}

type PathConfig struct {
	ProjectPath  string  `yaml:"projectPath"`
	JadxPath   	string  `yaml:"jadxPath"`
}

var (
	config *Config
	once   sync.Once
)

func LoadConfig(filename string) *Config {
	once.Do(func() {
		file, err := os.Open(filename)
		if err != nil {
			log.Fatalf("Error loading config file: %v", err)
		}
		defer file.Close()

		decoder := yaml.NewDecoder(file)
		if err := decoder.Decode(&config); err != nil {
			log.Fatalf("Error decoding config file: %v", err)
		}
	}) 

	return config
}