package dbo

import (
	"strings"

	"github.com/Cyberpull/gokit/fmt"
	"gorm.io/gorm"
)

type SeederHandler func(db *gorm.DB) (err error)

type DBSeeder interface {
	Seed(db *gorm.DB, entries []SeederEntry) (err error)
}

type SeederEntry interface {
	Name() string
	Handler(db *gorm.DB) (err error)
}

type dbSeeder struct {
	opts     *Options
	handlers []SeederHandler
}

func (x *dbSeeder) Add(handlers ...SeederHandler) {
	x.handlers = append(x.handlers, handlers...)
}

func (x *dbSeeder) Run(db *gorm.DB) (err error) {
	for _, handler := range x.handlers {
		tx := NewSession(db)

		if err = handler(tx); err != nil {
			return
		}
	}

	return
}

func (x *dbSeeder) RunEntry(db *gorm.DB, entry SeederEntry) (err error) {
	name := strings.TrimSpace(entry.Name())

	fmt.Magenta.Printf("Seeding '%v'... ", name)

	err = entry.Handler(db)

	if err != nil {
		return
	}

	fmt.Magenta.Println("Completed!")

	return
}

func (x *dbSeeder) Seed(db *gorm.DB, entries []SeederEntry) (err error) {
	for _, entry := range entries {
		tx := NewSession(db)

		if err = x.RunEntry(tx, entry); err != nil {
			return
		}
	}

	return
}

// ===================

var Seeder dbSeeder

func NewSeeder() DBSeeder {
	return newSeeder(Seeder.opts)
}

func newSeeder(opts *Options) *dbSeeder {
	s := &dbSeeder{opts: opts}
	initSeeders(s)
	return s
}

func initSeeders(x *dbSeeder) {
	x.handlers = []SeederHandler{}
}

func init() {
	initSeeders(&Seeder)
}
