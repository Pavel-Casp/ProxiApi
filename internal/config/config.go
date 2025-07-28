package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

type Config struct {
	API struct {
		BaseURL        string        `yaml:"base_url" envconfig:"API_BASE_URL" default:"https://jsonplaceholder.typicode.com"`
		RequestTimeout time.Duration `yaml:"request_timeout" envconfig:"REQUEST_TIMEOUT" default:"10s"`
	} `yaml:"api"`

	Retry struct {
		MaxRetries int           `yaml:"max_retries" envconfig:"MAX_RETRIES" default:"3"`
		WaitTime   time.Duration `yaml:"wait_time" envconfig:"RETRY_WAIT_TIME" default:"1s"`
		MaxWait    time.Duration `yaml:"max_wait" envconfig:"RETRY_MAX_WAIT" default:"5s"`
	} `yaml:"retry"`

	Debug struct {
		Enabled  bool   `yaml:"enabled" envconfig:"DEBUG_MODE" default:"false"`
		LogLevel string `yaml:"log_level" envconfig:"LOG_LEVEL" default:"info"`
	} `yaml:"debug"`
}

func Load() (*Config, error) {
	cfg := &Config{}

	// 1. Попробуем загрузить YAML-конфиг
	yamlPath := filepath.Join("configs", "config.yaml")
	if _, err := os.Stat(yamlPath); err == nil {
		yamlFile, err := os.ReadFile(yamlPath)
		if err != nil {
			return nil, fmt.Errorf("error reading YAML config: %w", err)
		}

		if err := yaml.Unmarshal(yamlFile, cfg); err != nil {
			return nil, fmt.Errorf("error parsing YAML config: %w", err)
		}
	} else {
		log.Println("YAML config not found, using defaults")
	}

	// 2. Загружаем .env файл (переопределяет YAML)
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			log.Printf("Warning: Error loading .env file: %v", err)
		}
	}

	// 3. Загружаем переменные окружения (переопределяют всё)
	if err := envconfig.Process("", cfg); err != nil {
		return nil, fmt.Errorf("failed to process env vars: %w", err)
	}

	return cfg, nil
}
