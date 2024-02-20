package firebase

import (
	"errors"

	"firebase.google.com/go/v4/auth"
)

// CreateUser creates a new user with the specified email and password
func (s *Service) CreateUser(email, password string) (*auth.UserRecord, error) {
	params := (&auth.UserToCreate{}).
		Email(email).
		Password(password)
	return s.auc.CreateUser(s.ctx, params)
}

// GetUser returns user record by UID, email, or phone number
func (s *Service) GetUser(identifier string, value interface{}) (*auth.UserRecord, error) {
	switch identifier {
	case "uid":
		if uid, ok := value.(string); ok {
			return s.auc.GetUser(s.ctx, uid)
		}
	case "email":
		if email, ok := value.(string); ok {
			return s.auc.GetUserByEmail(s.ctx, email)
		}
	case "phoneNumber":
		if phoneNumber, ok := value.(string); ok {
			return s.auc.GetUserByPhoneNumber(s.ctx, phoneNumber)
		}
	}

	return nil, errors.New("invalid identifier type or value")
}
