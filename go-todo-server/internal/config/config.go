package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env string `yaml:"env" env-default:"local"`

	Server struct {
		Port        int           `yaml:"port" env-default:"3001"`
		Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
		IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	} `yaml:"server"`

	Database struct {
		Driver   string `yaml:"driver" env-default:"postgres"`
		Host     string `yaml:"host"  env-default:"localhost"`
		Port     int    `yaml:"port" env-default:"5432"`
		User     string `yaml:"user" env:"DB_USER"`
		Password string `yaml:"password" env:"DB_PASSWORD"`
		Name     string `yaml:"name" env-default:"DB_NAME"`
	} `yaml:"database"`
}

func MustLoad() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load env")
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("failed to get config path")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file doesnt exist: %v", err)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}

	return &cfg
}
