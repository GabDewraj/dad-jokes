package config

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ErrCannotEstablishDbConnection = errors.New("cannot connect to db")

type Config struct {
	DbConfig     DbConfig
	Certificate  Certificate
	ApiKeyConfig ApiKeyConfig
}

type ApiKeyConfig struct {
	Header string
}
type DbConfig struct {
	Host     string
	User     string
	Password string
	Database string
	Port     int
}
type Certificate struct {
	CertificateFilePath string
	KeyFilePath         string
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
	appConfig.Certificate = newCertificate(env)
	return &appConfig, nil
}

func newCertificate(env string) Certificate {
	if env == "" {
		certificate := Certificate{
			CertificateFilePath: "./certificates/local.crt",
			KeyFilePath:         "./certificates/local.key",
		}
		return certificate
	}
	certificate := Certificate{
		CertificateFilePath: fmt.Sprintf("./certificates/%s.crt", env),
		KeyFilePath:         fmt.Sprintf("./certificates/%s.key", env),
	}
	return certificate
}

func NewDB(dbConfig *DbConfig, maxRetries int) (*gorm.DB, error) {
	// Connect to postgres in container
	path := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Database, dbConfig.Port)
	log.Print(dbConfig)

	var ConnectionRetryCount int
	var dbConn *gorm.DB
	err := ErrCannotEstablishDbConnection

	for err != nil {
		dbConn, err = gorm.Open(postgres.Open(path))
		if err != nil {
			ConnectionRetryCount += 1
			log.Print("Retry Attempt Number ", ConnectionRetryCount)
		}
		if ConnectionRetryCount == maxRetries {
			return nil, ErrCannotEstablishDbConnection
		}
		time.Sleep(1 * time.Second)
	}
	log.Print("Connection to the DB is successful")
	return dbConn, nil
}
