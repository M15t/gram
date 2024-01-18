package session

import (
	"himin-runar/internal/repo"
	"himin-runar/pkg/rbac"
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
