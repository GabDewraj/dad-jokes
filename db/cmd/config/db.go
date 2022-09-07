package config

import (
	"errors"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Joke struct {
	gorm.Model
	Body   string
	Author string
}

var ErrCannotEstablishDbConnection = errors.New("cannot connect to db")

var DummyData = []Joke{
	{
		Model:  gorm.Model{},
		Body:   "What did one pirate say to the other when he beat him at chess?<>Checkmatey.",
		Author: "github.com",
	},
	{
		Model:  gorm.Model{},
		Body:   "I burned 2000 calories today<>I left my food in the oven for too long.",
		Author: "github.com",
	},
	{
		Model:  gorm.Model{},
		Body:   "I broke my arm in two places. <>My doctor told me to stop going to those places.",
		Author: "github.com",
	},
	{
		Model:  gorm.Model{},
		Body:   "I quit my job at the coffee shop the other day. <>It was just the same old grind over and over.",
		Author: "github.com",
	},
	{
		Model:  gorm.Model{},
		Body:   "I visited a weight loss website...<>they told me I have to have cookies disabled.",
		Author: "github.com",
	},
	{
		Model:  gorm.Model{},
		Body:   "Did you hear about the famous Italian chef that recently died? <>He pasta way.",
		Author: "github.com",
	},
	{
		Model:  gorm.Model{},
		Body:   "Broken guitar for sale<>no strings attached.",
		Author: "github.com",
	},
	{
		Model:  gorm.Model{},
		Body:   "I could never be a plumber<>it’s too hard watching your life’s work go down the drain.",
		Author: "github.com",
	},
}

func NewDB(dbConfig *DbConfig, maxRetries int) (*gorm.DB, error) {
	// Connect to postgres in container
	path := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Database, dbConfig.Port)

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

func DropTables(dbConn *gorm.DB) error {
	// Drop Any Existing tables
	if err := dbConn.Migrator().DropTable(&Joke{}); err != nil {
		return err
	}
	log.Print("All previous Existing tables have been Dropped ...")
	return nil
}

func CreateTables(dbConn *gorm.DB) error {
	if err := dbConn.AutoMigrate(&Joke{}); err != nil {
		return err
	}
	log.Print("Tables Created Successfully")

	// Seed the Database with dummy Data
	tx := dbConn.Create(&DummyData)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
