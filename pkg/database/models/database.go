package models

import (
	"database/sql"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type Database struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Database string    // postgres, mysql, etc
	Name     string    // database name, e.g. data_warehouse_prod, user_service_dev
	URL      string    // database url, or connection string
}

func CheckPostgresURL(url string) error {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) BeforeCreate(tx *gorm.DB) (err error) {
	switch d.Database {
	case "postgres":
		err = CheckPostgresURL(d.URL)
		if err != nil {
			return err
		}
	default:
		return gorm.ErrInvalidTransaction
	}

	return nil
}

func (d *Database) BeforeSave(tx *gorm.DB) (err error) {
	switch d.Database {
	case "postgres":
		err = CheckPostgresURL(d.URL)
		if err != nil {
			return err
		}
	default:
		return gorm.ErrInvalidTransaction
	}

	return nil
}

func CreateURL(host, port, user, password, dbname string) string {
	url := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname

	return url
}

func GetActiveDatabases(db *gorm.DB) ([]Database, error) {
	var databases []Database

	dbs := db.Where("deleted_at IS NULL").Find(&databases)

	return databases, dbs.Error
}
