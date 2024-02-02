package memo

import (
	"net/http"

	"github.com/M15t/gram/pkg/server"
)

// Custom errors
var (
	ErrMemoNotFound = server.NewHTTPError(http.StatusBadRequest, "MEMO_NOTFOUND", "Memo not found")
)
