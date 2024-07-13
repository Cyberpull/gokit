package dbo

import (
	"cyberpull.com/gotk/v2/errors"

	"gorm.io/gorm"
)

type Instance interface {
	New() *gorm.DB
	DB(db ...*gorm.DB) (value *gorm.DB, err error)
	AddMigrations(models ...any)
	AddSeeders(handlers ...SeederHandler)
	Migrate(seed ...bool) (err error)
	Seed() (err error)
}

// ======================

type dbInstance struct {
	opts       *Options
	db         *gorm.DB
	migrations *dbMigration
	seeders    dbSeeder
}

func (s *dbInstance) New() *gorm.DB {
	return NewSession(s.db)
}

func (s *dbInstance) DB(db ...*gorm.DB) (value *gorm.DB, err error) {
	if len(db) > 0 && db[0] != nil {
		value = db[0]
		return
	}

	if s.db != nil {
		value = s.New()
		return
	}

	err = errors.New("Database connection not found")

	return
}

func (s *dbInstance) AddMigrations(models ...any) {
	s.migrations.Add(models...)
}

func (s *dbInstance) AddSeeders(handlers ...SeederHandler) {
	s.seeders.Add(handlers...)
}

func (s *dbInstance) Migrate(seed ...bool) (err error) {
	if err = s.migrations.Run(s.db); err != nil {
		return
	}

	if len(seed) > 0 && seed[0] {
		err = s.seeders.Run(s.db)
	}

	return
}

func (s *dbInstance) Seed() (err error) {
	return s.seeders.Run(s.db)
}

// ======================

func DB(i Instance, db ...*gorm.DB) (value *gorm.DB, err error) {
	if i == nil {
		err = errors.New("No database connection instance found.")
		return
	}

	return i.DB(db...)
}

func NewInstance(db *gorm.DB, opts *Options) Instance {
	return &dbInstance{
		db:         db,
		opts:       opts,
		migrations: newMigration(opts),
		seeders:    *newSeeder(opts),
	}
}
