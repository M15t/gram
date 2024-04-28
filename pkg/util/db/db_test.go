package dbutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMySQLDialector_OpenConnection(t *testing.T) {

	cfg := Config{
		Username: "testuser",
		Password: "testpassword",
		Host:     "localhost",
		Port:     3306,
		Database: "testdb",
		Params:   "charset=utf8mb4&parseTime=true&loc=UTC",
	}

	dialector := &MySQLDialector{}
	gormDialector, err := dialector.OpenConnection(cfg)

	assert.NoError(t, err)
	assert.NotNil(t, gormDialector)
}

func TestPostgreSQLDialector_OpenConnection(t *testing.T) {
	cfg := Config{
		Username: "testuser",
		Password: "testpass",
		Host:     "localhost",
		Port:     5432,
		Database: "testdb",
		Params:   "sslmode=require",
	}

	dialector := &PostgreSQLDialector{}
	gormDialector, err := dialector.OpenConnection(cfg)

	assert.NoError(t, err)
	assert.NotNil(t, gormDialector)
}

// Returns a MySQL dialector and no error when given valid configuration
func TestOpenConnection_ValidConfig_ReturnsDialectorAndNoError(t *testing.T) {
	cfg := Config{
		Username: "testuser",
		Password: "testpassword",
		Host:     "localhost",
		Port:     3306,
		Database: "testdb",
		Params:   "charset=utf8mb4&parseTime=true&loc=UTC",
	}

	dialector := &MySQLDialector{}
	gormDialector, err := dialector.OpenConnection(cfg)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if gormDialector == nil {
		t.Error("expected non-nil gorm.Dialector, but got nil")
	}
}
