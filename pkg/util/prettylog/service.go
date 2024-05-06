package prettylog

import (
	"bytes"
	"log/slog"
	"sync"
)

// OutputFormat custom type
type OutputFormat int

// custom
const (
	JSONFormat OutputFormat = iota
	TextFormat
)

// Handler represents a type that contains a slog.Handler, a function, a bytes.Buffer, and a sync.Mutex.
type Handler struct {
	h      slog.Handler
	r      func([]string, slog.Attr) slog.Attr
	b      *bytes.Buffer
	m      *sync.Mutex
	format OutputFormat
}

// NewHandler is a function that creates a new instance of the Handler struct
func NewHandler(opts *slog.HandlerOptions, format OutputFormat) *Handler {
	if opts == nil {
		opts = &slog.HandlerOptions{}
	}
	b := &bytes.Buffer{}

	var h slog.Handler
	switch format {
	case TextFormat:
		h = slog.NewTextHandler(b, opts)
	default:
		h = slog.NewJSONHandler(b, opts)
	}

	return &Handler{
		h:      h,
		r:      opts.ReplaceAttr,
		b:      b,
		m:      &sync.Mutex{},
		format: format,
	}
}
