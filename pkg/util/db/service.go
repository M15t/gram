package dbutil

import (
	"fmt"
	"net/url"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Config is the database configuration
type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
	Params   string
}

// New creates new database connection to the database server
func New(dialect string, cfg Config, gormConfig *gorm.Config) (db *gorm.DB, err error) {
	switch dialect {
	case "mysql":
		params, err := url.ParseQuery(cfg.Params)
		if err != nil {
			return nil, fmt.Errorf("invalid db params '%s': %w", cfg.Params, err)
		}
		if params.Get("charset") == "" {
			params.Set("charset", "utf8mb4")
		}
		if params.Get("parseTime") == "" {
			params.Set("parseTime", "true")
		}
		if params.Get("loc") == "" {
			params.Set("loc", "UTC")
		}

		// generate the connection string
		dbDsn := cfg.Username + ":" + cfg.Password + "@tcp(" + cfg.Host + ":" + fmt.Sprint(cfg.Port) + ")/" + cfg.Database + "?" + params.Encode()

		var datetimePrecision = 3
		db, err = gorm.Open(mysql.New(mysql.Config{
			DSN:                      dbDsn,
			DefaultStringSize:        255,
			DefaultDatetimePrecision: &datetimePrecision,
		}), gormConfig)
		if err != nil {
			return nil, err
		}
	case "postgres":
		params, err := url.ParseQuery(cfg.Params)
		if err != nil {
			return nil, fmt.Errorf("invalid db params '%s': %w", cfg.Params, err)
		}
		if params.Get("sslmode") == "" {
			params.Set("sslmode", "disable")
		}
		if params.Get("connect_timeout") == "" {
			params.Set("connect_timeout", "5")
		}

		// generate the connection string
		dbDsn := "postgres://" + cfg.Username + ":" + cfg.Password + "@" + cfg.Host + ":" + fmt.Sprint(cfg.Port) + "/" + cfg.Database + "?" + params.Encode()

		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN: dbDsn,
			// Note: set to false to disable implicit prepared statement usage, in case using pgbouncer for example
			PreferSimpleProtocol: true,
		}), gormConfig)
		if err != nil {
			return nil, err
		}
	case "sqlite3":
		var dbDsn string
		db, err = gorm.Open(sqlite.Open(dbDsn), gormConfig)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
