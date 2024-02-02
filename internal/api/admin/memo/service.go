package memo

import (
	"github.com/M15t/gram/internal/repo"
	"github.com/M15t/gram/pkg/rbac"
)

// New creates new memo application service
func New(repo *repo.Service, rbacSvc rbac.Intf) *Memo {
	return &Memo{repo: repo, rbac: rbacSvc}
}

// Memo represents memo application service
type Memo struct {
	repo *repo.Service
	rbac rbac.Intf
}
