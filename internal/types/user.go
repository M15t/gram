package types

import (
	"time"

	"github.com/M15t/gram/pkg/util/firebase"
)

// cosnt
const (
	UserStatusUnknown Status = iota
	UserStatusActive
	UserStatusBlocked
	UserStatausDeleted
)

// Status represents the status of user
type Status int

func (s Status) String() string {
	switch s {
	case UserStatusActive:
		return "active"
	case UserStatusBlocked:
		return "blocked"
	case UserStatausDeleted:
		return "deleted"
	default:
		return "unknown"
	}
}

// User represents the user model
// swagger:model
type User struct {
	Base
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`

	Password  string     `json:"-" gorm:"not null"`
	LastLogin *time.Time `json:"last_login,omitempty" gorm:"type:datetime(3)"`

	Phone           string     `json:"phone" gorm:"uniqueIndex:uix_users_phone"`
	PhoneVerifiedAt *time.Time `json:"phone_verified_at,omitempty" gorm:"type:datetime(3)"`
	OTP             *string    `json:"-" gorm:"varchar(10)"`
	OTPSentAt       *time.Time `json:"-" gorm:"type:datetime(3)"`
	Email           string     `json:"email" gorm:"uniqueIndex:uix_users_email"`
	EmailVerifiedAt *time.Time `json:"email_verified_at,omitempty" gorm:"type:datetime(3)"`

	Status string `json:"status" gorm:"type:varchar(20);default:active"` // active || blocked || deleted

	UID string `json:"uid" gorm:"type:varchar(36)"` // reflect to firebase UID
}

// UserFirebase represents the user model with firebase embed
// swagger:model
type UserFirebase struct {
	FirebaseEmbed *firebase.AuthenticatedUserInfo `json:"firebase_embed" gorm:"embedded"`
}
