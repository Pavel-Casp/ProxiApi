package config

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	APIBaseURL     string        `envconfig:"API_BASE_URL" default:"https://jsonplaceholder.typicode.com"`
	RequestTimeout time.Duration `envconfig:"REQUEST_TIMEOUT" default:"10s"`
	MaxRetries     int           `envconfig:"MAX_RETRIES" default:"3"`
	RetryWaitTime  time.Duration `envconfig:"RETRY_WAIT_TIME" default:"1s"`
	RetryMaxWait   time.Duration `envconfig:"RETRY_MAX_WAIT" default:"5s"`
	DebugMode      bool          `envconfig:"DEBUG_MODE" default:"false"`
}

func Load() *Config {
	// Загружаем .env файл
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, using default values")
	}

	var cfg Config
	err = envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("Failed to process env vars: %v", err)
	}

	return &cfg
}
