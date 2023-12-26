package repository

import (
	"context"
	"jeffersonroth/nangarekohara/domain"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type databaseRepository struct {
	db *gorm.DB
}

func NewDatabaseRepository(db *gorm.DB) domain.DatabaseRepository {
	return &databaseRepository{
		db: db,
	}
}

func (d *databaseRepository) Create(c context.Context, database domain.Database) error {
	err := d.db.Create(&database).Error
	if err != nil {
		return err
	}

	return nil
}

func (d *databaseRepository) Fetch(c context.Context) ([]domain.Database, error) {
	var databases []domain.Database

	err := d.db.Where("deleted_at IS NULL").Find(&databases).Error
	if err != nil {
		return nil, err
	}

	return databases, nil
}

func (d *databaseRepository) FetchByName(c context.Context, name string) ([]domain.Database, error) {
	var databases []domain.Database

	err := d.db.Where("name = ?", name).First(&databases).Error
	if err != nil {
		return nil, err
	}

	return databases, nil
}

func (d *databaseRepository) Update(c context.Context, database domain.Database) error {
	err := d.db.Save(&database).Error
	if err != nil {
		return err
	}

	return nil
}

func (d *databaseRepository) Delete(c context.Context, database domain.Database) error {
	err := d.db.Delete(&database).Error
	if err != nil {
		return err
	}

	return nil
}

func (d *databaseRepository) Validate(c context.Context, database domain.Database) error {
	switch database.Database {
	case "postgres":
		err := domain.CheckPostgresURL(database.URL)
		if err != nil {
			return err
		}
	default:
		return gorm.ErrInvalidTransaction
	}

	return nil
}
