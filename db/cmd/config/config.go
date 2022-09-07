package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DbConfig DbConfig
}
type DbConfig struct {
	Host     string
	User     string
	Password string
	Database string
	Port     int
}

func NewConfig(env string) (*Config, error) {
	viper.AddConfigPath("./config")
	viper.SetConfigType("yml")

	var appConfig Config

	if env == "" {
		viper.SetConfigName("terminal")
		log.Print("Running in the terminal ...")
	}

	if env != "" {
		viper.SetConfigName(env)
		log.Print("Running in container ...")
	}

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&appConfig); err != nil {
		return nil, err
	}

	return &appConfig, nil
}
