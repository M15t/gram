package auth

import (
	"database/sql"
	"gram/internal/types"
	"gram/pkg/server/middleware/jwt"
	"gram/pkg/util/ulidutil"
	"time"

	"github.com/labstack/echo/v4"
)

func (s *Auth) authenticate(c echo.Context, ai *AuthenticateInput) (*types.AuthToken, error) {
	// * noted: this case is support only 1 session per user
	// * multiple sessions per user is not supported yet need to be implemented later
	ctx := c.Request().Context()

	sessionID := ulidutil.NewString()
	accessTokenOutput := jwt.TokenOutput{}
	refreshTokenOutput := jwt.TokenOutput{}
	// * generate access token
	if err := s.jwt.GenerateToken(&jwt.TokenInput{
		Type: jwt.TypeTokenAccess,
		Claims: map[string]interface{}{
			"id":    ai.User.ID,
			"email": ai.User.Email,
			"name":  ai.User.FirstName + " " + ai.User.LastName,
			"role":  ai.User.Role,
		},
	}, &accessTokenOutput); err != nil {
		return nil, err
	}

	// * generate refresh token
	if err := s.jwt.GenerateToken(&jwt.TokenInput{
		Type: jwt.TypeTokenRefresh,
		Claims: map[string]interface{}{
			"id":  sessionID,
			"uid": ai.User.ID,
		},
	}, &refreshTokenOutput); err != nil {
		return nil, err
	}

	if ai.IsLogin {
		// * create session
		if err := s.repo.Session.Create(ctx, &types.Session{
			ID:        sessionID,
			UserID:    ai.User.ID,
			IPAddress: c.RealIP(),
			UserAgent: c.Request().UserAgent(),
			ExpiresAt: time.Now().Add(time.Duration(refreshTokenOutput.ExpiresIn) * time.Second),
			RefreshToken: sql.NullString{
				String: refreshTokenOutput.Token,
				Valid:  true,
			},
		}); err != nil {
			return nil, err
		}
	}

	// * update last_login
	if err := s.repo.User.Update(ctx, map[string]interface{}{
		"last_login": time.Now(),
	}, ai.User.ID); err != nil {
		return nil, err
	}

	// * update refresh_token
	if err := s.repo.Session.Update(ctx, map[string]interface{}{
		"refresh_token": sql.NullString{
			String: refreshTokenOutput.Token,
			Valid:  true,
		},
	}, sessionID); err != nil {
		return nil, err
	}

	// TODO: add more logic if needed

	return &types.AuthToken{
		AccessToken:  accessTokenOutput.Token,
		TokenType:    "bearer",
		ExpiresIn:    accessTokenOutput.ExpiresIn,
		RefreshToken: refreshTokenOutput.Token,
	}, nil
}
