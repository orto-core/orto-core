package store

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Database struct {
	DB  *gorm.DB
	dsn string
}

func NewDatabase(dsn string) *Database {
	return &Database{dsn: dsn}
}

func (db *Database) Connect() error {
	var err error
	db.DB, err = gorm.Open(postgres.Open(db.dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) Migrate(models ...interface{}) error {
	if db.DB == nil {
		return gorm.ErrInvalidDB
	}
	return db.DB.AutoMigrate(models...)
}

func InitDB(dsn string, models ...interface{}) error {
	db := NewDatabase(dsn)
	if err := db.Connect(); err != nil {
		log.Fatalf("Failed to connect to database, %v", err)
		return err
	}

	fmt.Println(models...)

	if err := db.Migrate(models...); err != nil {
		log.Fatalf("Failed to migrate models, %v", err)
		return err
	}

	DB = db.DB
	return nil
}
