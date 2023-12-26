package bootstrap

import (
	"context"
	"jeffersonroth/nangarekohara/pkg/database/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(env *Env) *gorm.DB {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPass
	dbName := env.DBName

	dsn := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " password=" + dbPass + " dbname=" + dbName + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Migrations
	db.AutoMigrate(&models.Database{})

	return db
}
