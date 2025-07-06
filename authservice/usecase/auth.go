package usecase

import "context"

type AuthUsecase interface {
	Login(ctx context.Context, input *LoginRequest) (output *LoginResponse, err error)
	RefreshToken(ctx context.Context, input *RefreshTokenRequest) (output *RefreshTokenResponse, err error)
}

type (
	LoginRequest struct {
		Identifier string `json:"identifier" binding:"required"`
		Password   string `json:"password" binding:"required,min=6"`
	}
	LoginResponse struct {
		ID           string  `json:"id"`
		Username     string  `json:"username"`
		RoleName     string  `json:"role"`
		Token        string  `json:"accessToken"`
		RefreshToken string  `json:"refreshToken"`
		ExpiresIn    float64 `json:"expires_in"`
		TokenType    string  `json:"token_type"`
	}
)

type (
	RefreshTokenRequest struct {
		Username string `json:"username"`
	}
	RefreshTokenResponse struct {
		Token     string  `json:"token"`
		ExpiresIn float64 `json:"expires_in"`
		TokenType string  `json:"token_type"`
	}
)
