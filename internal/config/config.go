package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

// Build information -ldflags .
const (
	version    string = "dev"
	commitHash string = "-"
)

var cfg *Config

// Project - contains all parameters project information.
type Project struct {
	Debug       bool   `yaml:"debug"`
	Name        string `yaml:"name"`
	Environment string `yaml:"environment"`
	Version     string
	CommitHash  string
}

// Database - contains all parameters database connection.
type Database struct {
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	Migrations string `yaml:"migrations"`
	Name       string `yaml:"name"`
	SslMode    string `yaml:"sslmode"`
	Driver     string `yaml:"driver"`
}

// Grpc - contains parameter address grpc.
type Grpc struct {
	Port              int    `yaml:"port"`
	MaxConnectionIdle int64  `yaml:"maxConnectionIdle"`
	Timeout           int64  `yaml:"timeout"`
	MaxConnectionAge  int64  `yaml:"maxConnectionAge"`
	Host              string `yaml:"host"`
}

// Rest - contains parameter rest json connection.
type Rest struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

// create struct Sentry wich contains configs for sentry
type Sentry struct {
	DSN string `yaml:"dsn"`
}

// Config - contains all configuration parameters in config package.
type Config struct {
	Project  Project  `yaml:"project"`
	Grpc     Grpc     `yaml:"grpc"`
	Rest     Rest     `yaml:"rest"`
	Database Database `yaml:"database"`
	//Metrics  Metrics  `yaml:"metrics"`
	//Jaeger   Jaeger   `yaml:"jaeger"`
	//Kafka    Kafka    `yaml:"kafka"`
	//Status   Status   `yaml:"status"`
	Sentry Sentry `yaml:"sentry"`
}



func ReadConfigYAML(filePath string) error {
	if cfg != nil {
		return nil
	}

	file, err := os.Open(filepath.Clean(filePath)
	if err != nil {
		return err
	}

	defer func() {
		_ = file.Close()
	}()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return err
	}

	cfg.Project.Version = version
	cfg.Project.CommitHash = commitHash
}
