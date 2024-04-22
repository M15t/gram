package slogger

import (
	"context"

	"gorm.io/gorm/logger"
)

type LogType int

const (
	Parallel LogType = iota
	Stdout
	JSON
	Smart
)

var gormCtxKey = struct{}{}

// WithContextGormLogger func
func WithContextGormLogger(ctx context.Context, val any) context.Context {
	return context.WithValue(ctx, gormCtxKey, val)
}

// GetContextGormLogger func
func GetContextGormLogger(ctx context.Context) any {
	return nil
}

type Log struct {
	l       logger.Interface
	logType LogType
}
