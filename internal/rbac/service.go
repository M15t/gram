package rbac

import (
	"github.com/M15t/gram/pkg/rbac"
)

// New returns new RBAC service
func New(enableLog bool) *rbac.RBAC {
	r := rbac.NewWithConfig(rbac.Config{EnableLog: enableLog})

	// Add permission for user role
	r.AddPolicy(RoleUser, ObjectUser, ActionReadAll)

	r.AddPolicy(RoleUser, ObjectMemo, ActionCreate)
	r.AddPolicy(RoleUser, ObjectMemo, ActionRead)
	r.AddPolicy(RoleUser, ObjectMemo, ActionReadAll)
	r.AddPolicy(RoleUser, ObjectMemo, ActionUpdate)
	r.AddPolicy(RoleUser, ObjectMemo, ActionDelete)

	// Add permission for admin role
	r.AddPolicy(RoleAdmin, ObjectUser, ActionAny)
	r.AddPolicy(RoleAdmin, ObjectSession, ActionAny)

	// Add permission for superadmin role
	r.AddPolicy(RoleSuperAdmin, ObjectAny, ActionAny)

	// Roles inheritance
	r.AddGroupingPolicy(RoleAdmin, RoleUser)
	r.AddGroupingPolicy(RoleSuperAdmin, RoleAdmin)

	r.GetModel().PrintPolicy()

	return r
}
