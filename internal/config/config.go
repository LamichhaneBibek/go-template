package config

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Password PasswordConfig
	JWT      JWTConfig
	Logger   LoggerConfig
}

type ServerConfig struct {
	InternalPort string
	ExternalPort string
	RunMode      string
}

type PostgresConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	DatabaseName    string
	SSLMode         string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifeTime time.Duration
}

type PasswordConfig struct {
	IncludeChars     bool
	IncludeNumbers   bool
	MinLength        int
	MaxLength        int
	IncludeUppercase bool
	IncludeLowercase bool
}

type JWTConfig struct {
	Secret                   string
	AccessExpirationMinutes  time.Duration
	RefreshExpirationMinutes time.Duration
	RefreshSecret            string
}

type LoggerConfig struct {
	FilePath string
	Encoding string
	Level    string
	Logger   string
}

func GetConfig() *Config {
	configPath := getConfigPath(os.Getenv("APP_ENV"))
	vipe, err := LoadConfig(configPath, "yml")
	if err != nil {
		log.Fatalf("Error in loading config %v", err)
	}
	config, err := ParseConfig(vipe)
	envPort := os.Getenv("PORT")
	if envPort != "" {
		config.Server.ExternalPort = envPort
		log.Printf("Set external port from environment -> %s", config.Server.ExternalPort)
	} else {
		config.Server.ExternalPort = config.Server.InternalPort
		log.Printf("Set external port from config -> %s", config.Server.ExternalPort)
	}
	if err != nil {
		log.Fatalf("Error in parsing config %v", err)
	}

	return config

}

func ParseConfig(vipe *viper.Viper) (*Config, error) {
	var config Config
	err := vipe.Unmarshal(&config)
	if err != nil {
		log.Printf("unable to parse config, %v", err)
		return nil, err
	}
	return &config, nil
}

func LoadConfig(configPath string, configType string) (*viper.Viper, error) {
	vipe := viper.New()
	vipe.SetConfigName(configPath)
	vipe.AddConfigPath(".")
	vipe.AutomaticEnv()
	err := vipe.ReadInConfig()
	if err != nil {
		log.Printf("unable to load config, %v", err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	return vipe, nil
}

func getConfigPath(env string) string {
	return "./internal/config/config"
}
