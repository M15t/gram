package secure

// When allowOrigins is not empty, the function should return a middleware function that adds CORS headers to the response for simple requests with allowed origins.
import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestSimpleCORSWithAllowedOrigins(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Set the request origin header
	req.Header.Set(echo.HeaderOrigin, "http://example.com")

	// Define the allowed origins
	allowOrigins := []string{"http://example.com"}

	// Call the SimpleCORS middleware
	middleware := SimpleCORS(allowOrigins)
	handler := middleware(func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// Perform the request
	err := handler(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "http://example.com", rec.Header().Get(echo.HeaderAccessControlAllowOrigin))
}
