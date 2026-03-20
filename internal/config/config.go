package config

import (
	"os"
	"runtime"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App struct {
		Name string `yaml:"name"`
		Port string `yaml:"port"`
	} `yaml:"app"`

	Providers struct {
		Gemini struct {
			APIKey string `yaml:"api_key"`
			Model  string `yaml:"model"`
		} `yaml:"gemini"`
	} `yaml:"providers"`
}

func Load() (*Config, error) {
	configPath := "configs/config.yaml"
	if runtime.GOOS == "windows" {
		configPath = "configs/config.windows.yaml"
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
