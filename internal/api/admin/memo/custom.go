package memo

import (
	"himin-runar/pkg/server"
	"net/http"
)

// Custom errors
var (
	ErrMemoNotFound = server.NewHTTPError(http.StatusBadRequest, "MEMO_NOTFOUND", "Memo not found")
)
