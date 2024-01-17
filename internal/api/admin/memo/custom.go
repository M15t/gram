package memo

import (
	"net/http"
	"runar-himmel/pkg/server"
)

// Custom errors
var (
	ErrMemoNotFound = server.NewHTTPError(http.StatusBadRequest, "MEMO_NOTFOUND", "Memo not found")
)
