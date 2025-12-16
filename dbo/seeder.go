package dbo

import (
	"strings"

	"github.com/Cyberpull/gokit/fmt"
	"gorm.io/gorm"
)

type SeederHandler func(db *gorm.DB) (err error)

type SeederEntry interface {
	Name() string
	Handler(db *gorm.DB) (err error)
}

type dbSeeder struct {
	opts     *Options
	entries  []SeederEntry
	handlers []SeederHandler
}

func (x *dbSeeder) Add(handlers ...SeederHandler) {
	x.handlers = append(x.handlers, handlers...)
}

func (x *dbSeeder) AddEntries(entries ...SeederEntry) {
	x.entries = append(x.entries, entries...)
}

func (x *dbSeeder) Run(db *gorm.DB) (err error) {
	for _, entry := range x.entries {
		tx := NewSession(db)

		if err = x.RunEntry(tx, entry); err != nil {
			return
		}
	}

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

	fmt.Magenta.Println(" Completed!")

	return
}

// ===================

var Seeder dbSeeder

func newSeeder(opts *Options) *dbSeeder {
	s := &dbSeeder{opts: opts}
	initSeeders(s)
	return s
}

func initSeeders(x *dbSeeder) {
	x.entries = []SeederEntry{}
	x.handlers = []SeederHandler{}
}

func init() {
	initSeeders(&Seeder)
}
