package usecase

import (
	"context"
	"time"

	"github.com/jeffersonroth/nangarekohara/domain"
)

type databaseUsecase struct {
	databaseRepository domain.DatabaseRepository
	contextTimeout     time.Duration
}

func NewDatabaseUsecase(d domain.DatabaseRepository, timeout time.Duration) domain.DatabaseUsecase {
	return &databaseUsecase{
		databaseRepository: d,
		contextTimeout:     timeout,
	}
}

func (d *databaseUsecase) Create(c context.Context, database domain.Database) error {
	ctx, cancel := context.WithTimeout(c, d.contextTimeout)
	defer cancel()

	return d.databaseRepository.Create(ctx, database)
}

func (d *databaseUsecase) Fetch(c context.Context) ([]domain.Database, error) {
	ctx, cancel := context.WithTimeout(c, d.contextTimeout)
	defer cancel()

	return d.databaseRepository.Fetch(ctx)
}

func (d *databaseUsecase) FetchByName(c context.Context, name string) ([]domain.Database, error) {
	ctx, cancel := context.WithTimeout(c, d.contextTimeout)
	defer cancel()

	return d.databaseRepository.FetchByName(ctx, name)
}

func (d *databaseUsecase) Update(c context.Context, database domain.Database) error {
	ctx, cancel := context.WithTimeout(c, d.contextTimeout)
	defer cancel()

	return d.databaseRepository.Update(ctx, database)
}

func (d *databaseUsecase) Delete(c context.Context, database domain.Database) error {
	ctx, cancel := context.WithTimeout(c, d.contextTimeout)
	defer cancel()

	return d.databaseRepository.Delete(ctx, database)
}

func (d *databaseUsecase) Validate(c context.Context, database domain.Database) error {
	ctx, cancel := context.WithTimeout(c, d.contextTimeout)
	defer cancel()

	return d.databaseRepository.Validate(ctx, database)
}
