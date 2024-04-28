package main

import (
	"embed"
	"log/slog"

	"github.com/M15t/gram/config"
	"github.com/M15t/gram/internal/api/root"
	"github.com/M15t/gram/internal/api/v1/admin/session"
	"github.com/M15t/gram/internal/api/v1/admin/user"
	"github.com/M15t/gram/internal/api/v1/auth"
	"github.com/M15t/gram/internal/db"
	"github.com/M15t/gram/internal/rbac"
	"github.com/M15t/gram/internal/repo"

	"github.com/M15t/gram/pkg/server"
	"github.com/M15t/gram/pkg/server/middleware/jwt"
	"github.com/M15t/gram/pkg/server/middleware/secure"
	"github.com/M15t/gram/pkg/server/middleware/slogger"
	"github.com/M15t/gram/pkg/util/crypter"
	"github.com/M15t/gram/pkg/util/prettylog"

	contextutil "github.com/M15t/gram/internal/api/context"

	"github.com/labstack/echo/v4"
)

// To embed SwaggerUI into api server using go:build tag
var (
	enableSwagger = false
	swaggerui     embed.FS
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	cfg, err := config.LoadAll()
	checkErr(err)

	// Create a slog logger, which:
	//   - Logs to stdout.
	logger := slog.New(prettylog.NewHandler(nil))
	filters := make([]slogger.Filter, 0)
	filters = append(filters, slogger.IgnorePathContains("swagger"))

	db, sqldb, err := db.New(cfg.DB, logger)
	checkErr(err)
	defer sqldb.Close()

	// Initialize HTTP server
	e := server.New(&server.Config{
		Port:              cfg.Server.Port,
		ReadHeaderTimeout: cfg.Server.ReadHeaderTimeout,
		ReadTimeout:       cfg.Server.ReadTimeout,
		WriteTimeout:      cfg.Server.WriteTimeout,
		AllowOrigins:      cfg.Server.AllowOrigins,
		Debug:             cfg.General.Debug,
	})

	e.Use(slogger.NewWithConfig(logger, slogger.Config{
		WithUserAgent:    true,
		WithRequestBody:  true,
		WithResponseBody: true,
		Filters:          filters,
	}))

	if enableSwagger {
		// Static page for SwaggerUI
		e.GET("/swagger-ui/*", echo.StaticDirectoryHandler(echo.MustSubFS(swaggerui, "swagger-ui"), false), secure.DisableCache())
	}

	// Initialize core services
	crypterSvc := crypter.New()
	repoSvc := repo.New(db)
	rbacSvc := rbac.New(cfg.General.Debug)
	jwtSvc := jwt.New(cfg.JWT.Algorithm, cfg.JWT.Secret, cfg.JWT.DurationAccessToken, cfg.JWT.DurationRefreshToken)

	// Initialize services
	authSvc := auth.New(repoSvc, jwtSvc, crypterSvc)
	sessionSvc := session.New(repoSvc, rbacSvc)
	userSvc := user.New(repoSvc, rbacSvc, crypterSvc)

	// Initialize root API
	root.NewHTTP(e)
	v1Router := e.Group("/v1")

	auth.NewHTTP(authSvc, v1Router.Group("/auth"))

	// Initialize admin APIs
	adminRouter := v1Router.Group("/admin")
	adminRouter.Use(jwtSvc.MWFunc(), contextutil.MWContext())
	session.NewHTTP(sessionSvc, adminRouter.Group("/sessions"))
	user.NewHTTP(userSvc, adminRouter.Group("/users"))

	server.Start(e, config.IsLambda())
}
