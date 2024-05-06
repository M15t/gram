package db

import (
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"github.com/imdatngo/gowhere"

	"github.com/M15t/gram/config"

	// _ "gorm.io/driver/sqlite" // DB adapter
	// _ "gorm.io/gorm/dialects/postgres" // DB adapter
	_ "gorm.io/driver/mysql" // DB adapter
	"gorm.io/gorm"

	dblogger "github.com/M15t/gram/pkg/util/db/logger"
	"github.com/M15t/gram/pkg/util/prettylog"

	// EnablePostgreSQL: remove the mysql package above, uncomment the following
	dbutil "github.com/M15t/gram/pkg/util/db"
)

// New creates new database connection to the database server
func New(cfg config.DB) (*gorm.DB, *sql.DB, error) {
	// Add your DB related stuffs here, such as:
	// - gorm.DefaultTableNameHandler
	// - gowhere.DefaultConfig

	// ! EnablePostgreSQL
	// gowhere.DefaultConfig.Dialect = gowhere.DialectPostgreSQL
	gowhere.DefaultConfig.Dialect = gowhere.DialectMySQL

	// ! logger config for slog-gorm
	handlerOptions := &slog.HandlerOptions{}
	switch cfg.Logging {
	case -4: // LevelDebug
		handlerOptions.Level = slog.LevelDebug
	case 0: // LevelInfo
		handlerOptions.Level = slog.LevelInfo
	case 4: // LevelWarn
		handlerOptions.Level = slog.LevelWarn
	case 8: // LevelError
		handlerOptions.Level = slog.LevelError
	}

	// slogger := slog.New(slog.NewTextHandler(os.Stdout, handlerOptions))
	slogger := slog.New(prettylog.NewHandler(nil, prettylog.TextFormat))

	gcfg := dblogger.NewConfig(slogger.Handler()).
		WithTraceAll(true).
		WithRequestID(true)

	myLogger := dblogger.NewWithConfig(gcfg)

	// ! default gorm logger
	// var myLogger logger.Interface
	// if cfg.Logging > 0 {
	// 	myLogger = logger.Default.LogMode(logger.LogLevel(cfg.Logging))
	// } else {
	// 	myLogger = logger.Discard
	// }

	// parse extra params, merge with default params
	// change to PostgreSQLDialector{} for PostgreSQL
	dbConn, err := dbutil.NewDBConnection(&dbutil.MySQLDialector{}, dbutil.Config{
		Username: cfg.Username,
		Password: cfg.Password,
		Host:     cfg.Host,
		Port:     cfg.Port,
		Database: cfg.Database,
		Params:   cfg.Params,
	})
	if err != nil {
		return nil, nil, err
	}

	// connect to db server
	db, err := gorm.Open(dbConn, &gorm.Config{
		Logger:                                   myLogger,
		AllowGlobalUpdate:                        false,
		CreateBatchSize:                          1000,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("cannot establish connection: %w", err)
	}

	// connection pool settings
	sqldb, err := db.DB()
	if err != nil {
		return nil, nil, fmt.Errorf("cannot get generic db instance: %w", err)
	}
	//! NOTE: These are not one-size-fits-all settings. Turn it based on your db settings!
	sqldb.SetMaxIdleConns(10)
	sqldb.SetMaxOpenConns(10)
	sqldb.SetConnMaxLifetime(30 * time.Minute)
	sqldb.SetConnMaxIdleTime(10 * time.Minute)

	return db, sqldb, nil
}
