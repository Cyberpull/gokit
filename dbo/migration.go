package dbo

import (
	"fmt"

	"gorm.io/gorm"
)

type dbMigration struct {
	opts   *Options
	models []any
}

func (m *dbMigration) Add(models ...any) {
	m.models = append(m.models, models...)
}

func (m *dbMigration) Run(db *gorm.DB, seed ...bool) (err error) {
	switch dbDriver(m.opts) {
	case DRIVER_MYSQL:
		db = db.Set("gorm:table_options", fmt.Sprintf(
			"ENGINE=%s CHARSET=%s COLLATE=%s",
			engine(m.opts),
			charset(m.opts),
			collation(m.opts),
		))
	}

	for _, model := range m.models {
		err = db.AutoMigrate(model)

		if err != nil {
			return
		}
	}

	if len(seed) > 0 && seed[0] {
		err = Seeder.Run(db)
	}

	return
}

// ======================

var Migration dbMigration

func newMigration(opts *Options) *dbMigration {
	m := &dbMigration{opts: opts}
	initMigration(m)
	return m
}

func initMigration(m *dbMigration) {
	m.models = make([]any, 0)
}

func init() {
	initMigration(&Migration)
}
