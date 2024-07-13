package dbo

import (
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

	switch dbDriver(opts) {
	case DRIVER_PGSQL:
		err = db.Exec(`SET DEFAULT_TRANSACTION_ISOLATION TO SERIALIZABLE`).Error
		// SET TRANSACTION ISOLATION LEVEL SERIALIZABLE;
		// SET DEFAULT_TRANSACTION_ISOLATION TO SERIALIZABLE;
	}

	i = NewInstance(db, opts)

	return
}
