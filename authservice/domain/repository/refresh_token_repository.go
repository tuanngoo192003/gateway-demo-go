package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/tuanngoo192003/gateway-demo-go/authservice/common/utils"
	"github.com/tuanngoo192003/gateway-demo-go/authservice/domain/ent"
	refreshTokenDomain "github.com/tuanngoo192003/gateway-demo-go/authservice/domain/ent/refreshtoken"
)

type RefreshTokenRepository interface {
	Create(ctx context.Context, input *ent.RefreshToken) error
	GetList(ctx context.Context) ([]*ent.RefreshToken, error)
	Update(ctx context.Context, input *ent.RefreshToken) (*ent.RefreshToken, error)
	Delete(ctx context.Context, id string) error
	DeleteByUsername(ctx context.Context, username string) error
	FindOneByFields(ctx context.Context, id string) (*ent.RefreshToken, error)
	FindManyByFields(ctx context.Context, fields map[string]any) ([]*ent.RefreshToken, error)
}

type RefreshTokenRepositoryImpl struct {
	client *ent.Client
}

func NewRefreshTokenRepository(client *ent.Client) *RefreshTokenRepositoryImpl {
	return &RefreshTokenRepositoryImpl{client: client}
}

// Create a new refresh token
func (r *RefreshTokenRepositoryImpl) Create(ctx context.Context, input *ent.RefreshToken) error {
	_, err := utils.Transaction(ctx, r.client, func(tx *ent.Tx) (*ent.RefreshToken, error) {
		return tx.RefreshToken.
			Create().
			SetID(input.ID).
			SetUsername(input.Username).
			SetToken(input.Token).
			SetExpiredAt(input.ExpiredAt).
			SetRoleName(input.RoleName).
			Save(ctx)
	})

	return err
}

// GetList returns all refresh tokens
func (r *RefreshTokenRepositoryImpl) GetList(ctx context.Context) ([]*ent.RefreshToken, error) {
	return r.client.RefreshToken.
		Query().
		All(ctx)
}

// Update a refresh token
func (r *RefreshTokenRepositoryImpl) Update(ctx context.Context, input *ent.RefreshToken) (*ent.RefreshToken, error) {
	if input.ID == uuid.Nil {
		return nil, errors.New("invalid refresh token ID")
	}

	return utils.Transaction(ctx, r.client, func(tx *ent.Tx) (*ent.RefreshToken, error) {

		return tx.RefreshToken.
			UpdateOneID(input.ID).
			SetUsername(input.Username).
			SetToken(input.Token).
			SetExpiredAt(input.ExpiredAt).
			SetRoleName(input.RoleName).
			Save(ctx)
	})
}

// Delete by UUID string
func (r *RefreshTokenRepositoryImpl) Delete(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	_, err = utils.Transaction(ctx, r.client, func(tx *ent.Tx) (any, error) {
		return nil, tx.RefreshToken.
			DeleteOneID(uid).
			Exec(ctx)
	})

	return err
}

// Delete by username
func (r *RefreshTokenRepositoryImpl) DeleteByUsername(ctx context.Context, username string) error {
	_, err := utils.Transaction(ctx, r.client, func(tx *ent.Tx) (any, error) {
		_, err := tx.RefreshToken.
			Delete().
			Where(refreshTokenDomain.UsernameEQ(username)).
			Exec(ctx)
		return nil, err
	})

	return err
}

// FindOneByFields by ID (you could extend this to use other fields too)
func (r *RefreshTokenRepositoryImpl) FindOneByFields(ctx context.Context, id string) (*ent.RefreshToken, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return r.client.RefreshToken.
		Get(ctx, uid)
}

// FindManyByFields using dynamic field matching
func (r *RefreshTokenRepositoryImpl) FindManyByFields(ctx context.Context, fields map[string]any) ([]*ent.RefreshToken, error) {
	query := r.client.RefreshToken.Query()

	for key, value := range fields {
		switch key {
		case "username":
			if v, ok := value.(string); ok {
				query = query.Where(refreshTokenDomain.UsernameEQ(v))
			}
		case "token":
			if v, ok := value.(string); ok {
				query = query.Where(refreshTokenDomain.TokenEQ(v))
			}
		case "roleName":
			if v, ok := value.(refreshTokenDomain.RoleName); ok {
				query = query.Where(refreshTokenDomain.RoleNameEQ(v))
			}
		case "expiredAt":
			if v, ok := value.(string); ok {
				query = query.Where(refreshTokenDomain.ExpiredAtEQ(v))
			}
		}
	}

	return query.All(ctx)
}
