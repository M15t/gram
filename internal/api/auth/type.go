package auth

import "himin-runar/internal/types"

// Credentials represents login request data
// swagger:model
type Credentials struct {
	// example: loki@himin-runar.sky
	Email string `json:"email" form:"email" validate:"required_without=Username"`
	// example: user123!@#
	Password string `json:"password" form:"password" validate:"required"`

	// This is for SwaggerUI authentication which only support `username` field
	// swagger:ignore
	Username string `json:"username" form:"username"`
	// example: app
	GrantType string `json:"grant_type" form:"grant_type" validate:"required"`
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
