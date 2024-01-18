package session

import (
	"himin-runar/pkg/server"
	"net/http"
)

// Custom errors
var (
	ErrSessionNotFound = server.NewHTTPError(http.StatusBadRequest, "SESSION_NOTFOUND", "Session not found")
)
