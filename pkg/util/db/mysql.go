package dbutil

import (
	"fmt"
	"net/url"

	"github.com/M15t/gram/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySQLDialector is the MySQL database connection
type MySQLDialector struct{}

// OpenConnection opens a new MySQL connection
func (d *MySQLDialector) OpenConnection(cfg config.DB) (gorm.Dialector, error) {
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
	dbConn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		params.Encode(),
	)

	var datetimePrecision = 3
	return mysql.New(mysql.Config{
		DSN:                      dbConn,
		DefaultStringSize:        255,
		DefaultDatetimePrecision: &datetimePrecision,
	}), nil
}
