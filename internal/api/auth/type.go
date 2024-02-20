package auth

import "github.com/M15t/gram/internal/types"

// Credentials represents login request data
// swagger:model
type Credentials struct {
	// example: loki@gram.sky
	Email string `json:"email" form:"email" validate:"required_without=Username"`
	// example: user123!@#
	Password string `json:"password" form:"password"`

	// This is for SwaggerUI authentication which only support `username` field
	// swagger:ignore
	Username string `json:"username" form:"username"`
	// example: app
	GrantType string `json:"grant_type" form:"grant_type" validate:"required"`
}

// RegisterData represents register request data
// swagger:model
type RegisterData struct {
	// example:
	FirstName string `json:"first_name" form:"first_name" validate:"required"`
	// example:
	LastName string `json:"last_name" form:"last_name" validate:"required"`
	// example: user
	Role string `json:"role" form:"role" validate:"required"`
}

// RefreshTokenData represents refresh token request data
// swagger:model
type RefreshTokenData struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

// AuthenticateInput represents internal authenticate data
type AuthenticateInput struct {
	User    *types.User
	IsLogin bool
}
