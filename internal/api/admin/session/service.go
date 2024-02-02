package session

import (
	"github.com/M15t/gram/internal/repo"
	"github.com/M15t/gram/pkg/rbac"
)

// New creates new session application service
func New(repo *repo.Service, rbacSvc rbac.Intf) *Session {
	return &Session{repo: repo, rbac: rbacSvc}
}

// Session represents session application service
type Session struct {
	repo *repo.Service
	rbac rbac.Intf
}
