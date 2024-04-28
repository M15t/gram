package dbutil

import (
	"context"

	"github.com/M15t/gram/pkg/util/threadsafe"
)

// LogType defines log type
type LogType int

// custom
const (
	Parallel LogType = iota
	Stdout
	JSON
	Smart
)

var gormCtxKey = struct{}{}

// WithContextGormLogger takes a context and threadsafe slice (string) as inputs and returns a new context with a value
func WithContextGormLogger(ctx context.Context, w *threadsafe.SimpleSafeSlice[string]) context.Context {
	return context.WithValue(ctx, gormCtxKey, w)
}

// GetContextGormLogger takes a context as input and returns a pointer to a SimpleSafeSlice of strings
func GetContextGormLogger(ctx context.Context) *threadsafe.SimpleSafeSlice[string] {
	if w, ok := ctx.Value(gormCtxKey).(*threadsafe.SimpleSafeSlice[string]); ok {
		return w
	}
	return nil
}
