package dbo

import (
	"time"

	"gorm.io/gorm"
)

func Connect(opts *Options) (i Instance, err error) {
	var db *gorm.DB
	var conn gorm.Dialector

	config := opts.getConfig()

	if conn, err = dialector(opts); err != nil {
		return
	}

	if db, err = gorm.Open(conn, config); err != nil {
		return
	}

	sqldb, err := db.DB()

	if err != nil {
		return
	}

	if opts.MaxOpenConns > 0 {
		sqldb.SetMaxOpenConns(opts.MaxOpenConns)
	}

	if opts.MaxIdleConns > 0 {
		sqldb.SetMaxIdleConns(opts.MaxIdleConns)
	}

	if opts.ConnMaxLifetime > 0 {
		duration := time.Millisecond * time.Duration(opts.ConnMaxLifetime)
		sqldb.SetConnMaxLifetime(duration)
	}

	if opts.ConnMaxIdleTime > 0 {
		duration := time.Millisecond * time.Duration(opts.ConnMaxIdleTime)
		sqldb.SetConnMaxIdleTime(duration)
	}

	db.Use(NewPlugin())

	switch dbDriver(opts) {
	case DRIVER_PGSQL:
		err = db.Exec(`SET DEFAULT_TRANSACTION_ISOLATION TO SERIALIZABLE`).Error
		// SET TRANSACTION ISOLATION LEVEL SERIALIZABLE;
		// SET DEFAULT_TRANSACTION_ISOLATION TO SERIALIZABLE;
	}

	i = NewInstance(db, opts)

	return
}
