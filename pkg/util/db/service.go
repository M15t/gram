package dbutil

import (
	"gorm.io/gorm"
)

// Config stores configurations
type Config struct {
	Username, Password, Host, Params, Database string
	Port                                       int
}

// DBConnector represents an interface for establishing database connections.
type DBConnector interface {
	OpenConnection(cfg Config) (gorm.Dialector, error)
}

// NewDBConnection establishes a new database connection.
func NewDBConnection(connector DBConnector, cfg Config) (gorm.Dialector, error) {
	return connector.OpenConnection(cfg)
}
