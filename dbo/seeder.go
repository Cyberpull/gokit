package dbo

import "gorm.io/gorm"

type SeederHandler func(db *gorm.DB) (err error)

type dbSeeder struct {
	opts     *Options
	handlers []SeederHandler
}

func (s *dbSeeder) Add(handlers ...SeederHandler) {
	s.handlers = append(s.handlers, handlers...)
}

func (s *dbSeeder) Run(db *gorm.DB) (err error) {
	for _, handler := range s.handlers {
		tx := NewSession(db)

		if err = handler(tx); err != nil {
			return
		}
	}

	return
}

// ===================

var Seeder dbSeeder

func newSeeder(opts *Options) *dbSeeder {
	s := &dbSeeder{opts: opts}
	initSeeders(s)
	return s
}

func initSeeders(s *dbSeeder) {
	s.handlers = make([]SeederHandler, 0)
}

func init() {
	initSeeders(&Seeder)
}
