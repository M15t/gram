package secure

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// CORSConfig represents secure specific CORS config
type CORSConfig struct {
	AllowOrigins []string
	AllowMethods []string
}

// Headers adds general security headers for basic security measures
func Headers(securityPolicy string) echo.MiddlewareFunc {
	return middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "DENY",
		HSTSMaxAge:            31536000,
		HSTSExcludeSubdomains: true,
		ContentSecurityPolicy: securityPolicy,
	})
}

// CORS adds Cross-Origin Resource Sharing support
func CORS(cfg CORSConfig) echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:  cfg.AllowOrigins,
		AllowMethods:  cfg.AllowMethods,
		AllowHeaders:  []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders: []string{"Etag"},
		MaxAge:        86400,
	})
}

// DisableCache sets the Cache-Control directive to no-store.
func DisableCache() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			c.Response().Header().Set("Cache-Control", "no-store")
			return next(c)
		}
	}
}

// SimpleCORS returns a CORS middleware with minimum configurations. Preflighted request is not allowed though.
func SimpleCORS(allowOrigins []string) echo.MiddlewareFunc {
	if len(allowOrigins) == 0 {
		allowOrigins = []string{"*"}
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()

			// Check allowed origins
			origin := req.Header.Get(echo.HeaderOrigin)
			allowed := ""
			for _, o := range allowOrigins {
				if o == "*" || o == origin {
					allowed = o
					break
				}
			}

			// Simple request
			switch req.Method {
			case http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete:
				res.Header().Add(echo.HeaderVary, echo.HeaderOrigin)
				res.Header().Set(echo.HeaderAccessControlAllowOrigin, allowed)
				return next(c)
			}

			// Preflight request is only allowed when "all" origins are allowed
			if req.Method == http.MethodOptions && allowed == "*" {
				res.Header().Add(echo.HeaderVary, echo.HeaderOrigin)
				res.Header().Add(echo.HeaderVary, echo.HeaderAccessControlRequestMethod)
				res.Header().Add(echo.HeaderVary, echo.HeaderAccessControlRequestHeaders)
				res.Header().Set(echo.HeaderAccessControlAllowOrigin, "*")
				res.Header().Set(echo.HeaderAccessControlAllowMethods, "*")
				res.Header().Set(echo.HeaderAccessControlAllowHeaders, "*")
				return c.NoContent(http.StatusNoContent)
			}

			return echo.ErrMethodNotAllowed
		}
	}
}
