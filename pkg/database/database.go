package database

import (
	"os"

	"github.com/jeffersonroth/nangarekohara/pkg/database/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB       *gorm.DB
	Host     = os.Getenv("DB_HOST")
	Port     = os.Getenv("DB_PORT")
	User     = os.Getenv("DB_USER")
	Password = os.Getenv("DB_PASSWORD")
	Dbname   = os.Getenv("DB_NAME")
)

func InitDB() (*gorm.DB, error) {
	dsn := "host=" + Host + " port=" + Port + " user=" + User + " password=" + Password + " dbname=" + Dbname + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// Migrations
	db.AutoMigrate(&models.Database{})

	return db, err
}
