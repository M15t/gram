package dbutil

import (
	"fmt"
	"net/url"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgreSQLDialector is the PostgreSQL database connection
type PostgreSQLDialector struct{}

// OpenConnection opens a new PostgreSQL connection
func (d *PostgreSQLDialector) OpenConnection(cfg Config) (gorm.Dialector, error) {
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
	dbConn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		params.Encode(),
	)

	return postgres.New(postgres.Config{
		DSN: dbConn,
		// Note: set to false to disable implicit prepared statement usage, in case using pgbouncer for example
		PreferSimpleProtocol: true,
	}), nil
}
