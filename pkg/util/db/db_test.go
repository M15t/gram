package dbutil_test

import (
	"testing"

	"github.com/M15t/gram/config"
	dbutil "github.com/M15t/gram/pkg/util/db"
	"github.com/stretchr/testify/assert"
)

func TestMySQLDialector_OpenConnection(t *testing.T) {

	cfg := config.DB{
		Username: "testuser",
		Password: "testpassword",
		Host:     "localhost",
		Port:     3306,
		Database: "testdb",
		Params:   "charset=utf8mb4&parseTime=true&loc=UTC",
	}

	dialector := &dbutil.MySQLDialector{}
	gormDialector, err := dialector.OpenConnection(cfg)

	assert.NoError(t, err)
	assert.NotNil(t, gormDialector)
}

func TestPostgreSQLDialector_OpenConnection(t *testing.T) {
	cfg := config.DB{
		Username: "testuser",
		Password: "testpass",
		Host:     "localhost",
		Port:     5432,
		Database: "testdb",
		Params:   "sslmode=require",
	}

	dialector := &dbutil.PostgreSQLDialector{}
	gormDialector, err := dialector.OpenConnection(cfg)

	assert.NoError(t, err)
	assert.NotNil(t, gormDialector)
}
