package config

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig
	Redis RedisConfig
	Postgres PostgresConfig
}

type ServerConfig struct {
	Port    string
	RunMode string
}

type RedisConfig struct {
	Host               string
	Port               string
	Password           string
	Db                 string
	DialTimeout        time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	PoolSize           int
	PoolTimeout        time.Duration
	IdleCheckFrequency time.Duration
}

type PostgresConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	DbName          string
	SSLMode         string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

func GetConfig() *Config {
	fileName := getConfigPath(os.Getenv("APP_ENV"))
	v, err := loadConfig(fileName, "yml")
	if err != nil {
		log.Fatalf("Unable to load config %v", err)
	}
	cfg, err := parseConfig(v)
	if err != nil {
		log.Fatalf("Unable to parse config %v", err)
	}
	return cfg
}

func parseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func loadConfig(fileName, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigName(fileName)
	v.AddConfigPath("../")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		log.Printf("can not read config")
		return nil, err
	}
	return v, nil
}

func getConfigPath(env string) string {
	if env == "docker" {
		return "/config/config-docker"
	} else if env == "production" {
		return "/config/config-production"
	} else {
		return "/config/config-development"
	}
}
