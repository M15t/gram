package dbutil

import (
	"github.com/M15t/gram/config"
	"gorm.io/gorm"
)

// DBConnector represents an interface for establishing database connections.
type DBConnector interface {
	OpenConnection(cfg config.DB) (gorm.Dialector, error)
}

// NewDBConnection establishes a new database connection.
func NewDBConnection(connector DBConnector, cfg config.DB) (gorm.Dialector, error) {
	return connector.OpenConnection(cfg)
}
