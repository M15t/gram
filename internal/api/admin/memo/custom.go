package memo

import (
	"gram/pkg/server"
	"net/http"
)

// Custom errors
var (
	ErrMemoNotFound = server.NewHTTPError(http.StatusBadRequest, "MEMO_NOTFOUND", "Memo not found")
)
