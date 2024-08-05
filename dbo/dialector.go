package dbo

import (
	"fmt"

	"github.com/Cyberpull/gokit/errors"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func dialector(opts *Options) (conn gorm.Dialector, err error) {
	switch dbDriver(opts) {
	case DRIVER_MYSQL:
		conn = func() gorm.Dialector {
			if opts.DSN != "" {
				return mysql.Open(opts.DSN)
			}

			return mysql.Open(fmt.Sprintf(
				"%s:%s@tcp(%s:%s)/%s?charset=%s&collation=%s&parseTime=True&loc=Local",
				opts.Username,
				opts.Password,
				opts.Host,
				opts.Port,
				opts.DBName,
				charset(opts),
				collation(opts),
			))
		}()

	case DRIVER_PGSQL:
		conn = func() gorm.Dialector {
			if opts.DSN != "" {
				return postgres.Open(opts.DSN)
			}

			return postgres.Open(fmt.Sprintf(
				"postgres://%s:%s@%s:%s/%s",
				opts.Username,
				opts.Password,
				opts.Host,
				opts.Port,
				opts.DBName,
			))
		}()

	case DRIVER_SQLITE:
		conn = sqlite.Open(opts.DSN)

	default:
		err = errors.New("DB Driver not available")
	}

	return
}

func dbDriver(opts *Options) DRIVER {
	if opts == nil {
		return ""
	}

	if opts.Driver == "" {
		opts.Driver = DRIVER_PGSQL
	}

	return opts.Driver
}

func engine(opts *Options) ENGINE {
	if opts.Engine == "" {
		opts.Engine = ENGINE_INNODB
	}

	return opts.Engine
}

func charset(opts *Options) string {
	if opts.Charset == "" {
		opts.Charset = "utf8mb4"
	}

	return opts.Charset
}

func collation(opts *Options) string {
	if opts.Collation == "" {
		opts.Collation = "utf8mb4_general_ci"
	}

	return opts.Collation
}
