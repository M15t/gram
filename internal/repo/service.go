package repo

import "gorm.io/gorm"

// Service provides all databases
type Service struct {
	User    *User
	Session *Session
	Memo    *Memo
}

// New creates db service
func New(db *gorm.DB) *Service {
	return &Service{
		User:    NewUser(db),
		Session: NewSession(db),
		Memo:    NewMemo(db),
	}
}
