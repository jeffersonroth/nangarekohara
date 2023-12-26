package domain

import (
	"context"
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

type DatabaseRepository interface {
	Create(c context.Context, database Database) error
	Fetch(c context.Context) ([]Database, error)
	FetchByName(c context.Context, name string) ([]Database, error)
	Update(c context.Context, database Database) error
	Delete(c context.Context, database Database) error
	Validate(c context.Context, database Database) error
}

type DatabaseUsecase interface {
	Create(c context.Context, database Database) error
	Fetch(c context.Context) ([]Database, error)
	FetchByName(c context.Context, name string) ([]Database, error)
	Update(c context.Context, database Database) error
	Delete(c context.Context, database Database) error
	Validate(c context.Context, database Database) error
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

	err = CheckUniqueName(tx, d.Name)
	if err != nil {
		return err
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

	err = CheckUniqueName(tx, d.Name)
	if err != nil {
		return err
	}

	return nil
}

func CreateURL(host, port, user, password, dbname string) string {
	url := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname

	return url
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

func CheckUniqueName(db *gorm.DB, name string) error {
	var database Database
	err := db.Where("name = ?", name).First(&database).Error
	if err != nil {
		return err
	}

	return nil
}
