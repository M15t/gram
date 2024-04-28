package repo

import (
	"context"

	"github.com/M15t/gram/internal/types"

	repoutil "github.com/M15t/gram/pkg/util/repo"

	"gorm.io/gorm"
)

// Session represents the client for session table
type Session struct {
	*repoutil.Repo[types.Session]
}

// NewSession returns a new session database instance
func NewSession(gdb *gorm.DB) *Session {
	return &Session{repoutil.NewRepo[types.Session](gdb)}
}

// FindByID finds a session by the given ID and preload User
func (r *Session) FindByID(ctx context.Context, id, userID string) (*types.Session, error) {
	rec := &types.Session{}
	if err := r.GDB.WithContext(ctx).Preload(`User`).Where(`id = ? AND user_id = ? AND is_blocked = false`, id, userID).Take(rec).Error; err != nil {
		return nil, err
	}

	return rec, nil
}

// DeleteExpired deletes expired sessions
func (r *Session) DeleteExpired(ctx context.Context, userID string) error {
	return r.GDB.Delete(&types.Session{}, `expires_at < NOW() AND user_id = ?`, userID).Error
}
