package session

import (
	"net/http"

	"github.com/M15t/gram/pkg/server"
)

// Custom errors
var (
	ErrSessionNotFound = server.NewHTTPError(http.StatusBadRequest, "SESSION_NOTFOUND", "Session not found")
)
