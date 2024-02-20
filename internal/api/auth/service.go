package auth

import (
	"firebase.google.com/go/v4/auth"
	"github.com/M15t/gram/internal/repo"
	"github.com/M15t/gram/pkg/server/middleware/jwt"

	gjwt "github.com/golang-jwt/jwt/v5"
)

// New creates new auth service
func New(repo *repo.Service, jwt JWT, cr Crypter, fb Firebase) *Auth {
	return &Auth{
		repo: repo,
		jwt:  jwt,
		cr:   cr,
		fb:   fb,
	}
}

// Auth represents auth application service
type Auth struct {
	repo *repo.Service
	jwt  JWT
	cr   Crypter
	fb   Firebase
}

// JWT represents token generator (jwt) interface
type JWT interface {
	GenerateToken(*jwt.TokenInput, *jwt.TokenOutput) error
	ParseToken(string) (*gjwt.Token, error)
}

// Crypter represents security interface
type Crypter interface {
	CompareHashAndPassword(string, string) bool
}

// Firebase represents firebase interface
type Firebase interface {
	CreateUser(string, string) (*auth.UserRecord, error)
	GetUser(identifier string, value interface{}) (*auth.UserRecord, error)
}
