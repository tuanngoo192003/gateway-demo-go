package interactor

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/tuanngoo192003/gateway-demo-go/authservice/common/utils"
	"github.com/tuanngoo192003/gateway-demo-go/authservice/domain/ent"
	"github.com/tuanngoo192003/gateway-demo-go/authservice/domain/repository"
	"github.com/tuanngoo192003/gateway-demo-go/authservice/infra/config"
	"github.com/tuanngoo192003/gateway-demo-go/authservice/usecase"

	refreshTokenDomain "github.com/tuanngoo192003/gateway-demo-go/authservice/domain/ent/refreshtoken"
)

type AuthInteractor struct {
	repository      *repository.AccountRepositoryImpl
	tokenRepository *repository.RefreshTokenRepositoryImpl
	jwtSecret       []byte
	tokenExpiration time.Duration
}

func NewAuthInteractor(repository *repository.AccountRepositoryImpl, tokenRepository *repository.RefreshTokenRepositoryImpl, jwtSecret string) *AuthInteractor {
	return &AuthInteractor{
		repository:      repository,
		jwtSecret:       []byte(jwtSecret),
		tokenExpiration: time.Hour * 24,
	}
}

func (ai *AuthInteractor) Login(ctx context.Context, input *usecase.LoginRequest) (output *usecase.LoginResponse, err error) {
	log := config.GetLogger()

	account, err := getCredential(ctx, input, ai)
	if err != nil {
		return nil, err
	}
	if !utils.CheckPasswordHash(input.Password, account.Password) {
		log.Error("Invalid credential")
		return &usecase.LoginResponse{}, errors.New("invalid credential")
	}

	/* Generate JWT token with claims */
	now := time.Now()
	jwt.WithSubject(account.Username)
	claims := jwt.MapClaims{
		"username": account.Username,
		"roles":    []interface{}{account.Edges.Role.RoleName},
		"iat":      now.Unix(),
		"exp":      now.Add(ai.tokenExpiration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(ai.jwtSecret)
	if err != nil {
		log.Error("Token generation failed")
		return &usecase.LoginResponse{}, errors.New("token generation failed")
	}

	refreshTokenString := uuid.New().String()
	ai.saveRefreshToken(ctx, refreshTokenString, account.Username, string(account.Edges.Role.RoleName))

	return &usecase.LoginResponse{
		ID:           account.ID.String(),
		Username:     account.Username,
		RoleName:     string(account.Edges.Role.RoleName),
		Token:        tokenString,
		RefreshToken: refreshTokenString,
		ExpiresIn:    ai.tokenExpiration.Seconds(),
		TokenType:    "Bearer",
	}, nil
}

func (ai *AuthInteractor) saveRefreshToken(ctx context.Context, refreshToken string, username string, roleName string) (*ent.RefreshToken, error) {
	log := config.GetLogger()

	if err := ai.tokenRepository.DeleteByUsername(ctx, username); err != nil {
		log.Error("Failed to delete old refresh tokens: " + err.Error())
		return nil, err

	}

	// Insert new token
	expiredAt := time.Now().Add(time.Hour).Format("2006-01-02 15:04:05")

	parsedRoleName, err := ParseRoleName(roleName)
	if err != nil {
		log.Errorf("Failed to parsedRoleName: %v", err.Error())
	}
	err = ai.tokenRepository.Create(ctx, &ent.RefreshToken{
		Username:  username,
		Token:     refreshToken,
		ExpiredAt: expiredAt,
		RoleName:  parsedRoleName,
	})
	if err != nil {
		log.Error("Failed to save new refresh token: " + err.Error())
		return nil, err
	}
	return nil, nil
}

func ParseRoleName(s string) (refreshTokenDomain.RoleName, error) {
	switch s {
	case string(refreshTokenDomain.RoleNameADMIN.String()):
		return refreshTokenDomain.RoleNameADMIN, nil
	case string(refreshTokenDomain.RoleNameMANAGER.String()):
		return refreshTokenDomain.RoleNameMANAGER, nil
	case string(refreshTokenDomain.RoleNameCUSTOMER.String()):
		return refreshTokenDomain.RoleNameCUSTOMER, nil
	default:
		return "", fmt.Errorf("invalid role name: %s", s)
	}
}

func getCredential(ctx context.Context, input *usecase.LoginRequest, ai *AuthInteractor) (account *ent.Account, err error) {
	var accounts []*ent.Account

	if utils.IsEmail(account.Email) {
		accounts, err = ai.repository.FindManyByFields(ctx, map[string]any{"username": input.Identifier})
	} else {
		accounts, err = ai.repository.FindManyByFields(ctx, map[string]any{"username": input.Identifier})
	}

	if err == sql.ErrNoRows {
		return &ent.Account{}, err
	}

	if err != nil {
		return &ent.Account{}, err
	}

	account = accounts[0]
	return
}

func (ai *AuthInteractor) RefreshToken(ctx context.Context, input *usecase.RefreshTokenRequest) (output *usecase.RefreshTokenResponse, err error) {
	log := config.GetLogger()
	now := time.Now()

	var token *ent.RefreshToken
	tokens, err := ai.tokenRepository.FindManyByFields(ctx, map[string]any{"username": input.Username})
	if err != nil {
		log.Error(err.Error())
		return &usecase.RefreshTokenResponse{}, err
	}
	token = tokens[0]

	jwt.WithSubject(token.Username)
	claims := jwt.MapClaims{
		"roles": []interface{}{token.RoleName},
		"iat":   now.Unix(),
		"exp":   now.Add(ai.tokenExpiration).Unix(),
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := newToken.SignedString(ai.jwtSecret)
	if err != nil {
		log.Error("Token generation failed")
		return &usecase.RefreshTokenResponse{}, errors.New("Token generation failed")
	}

	return &usecase.RefreshTokenResponse{
		Token:     tokenString,
		ExpiresIn: ai.tokenExpiration.Seconds(),
		TokenType: "Bearer",
	}, nil
}
