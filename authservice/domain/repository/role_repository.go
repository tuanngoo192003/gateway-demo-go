package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/tuanngoo192003/gateway-demo-go/authservice/domain/ent"
)

type RoleRepository interface {
	Create(ctx context.Context, input *ent.Role) error
	GetList(ctx context.Context) ([]*ent.Role, error)
	Update(ctx context.Context, input *ent.Role) (*ent.Role, error)
	Delete(ctx context.Context, id string) error
}

type RoleRepositoryImpl struct {
	client *ent.Client
}

func NewRoleRepository(client *ent.Client) *RoleRepositoryImpl {
	return &RoleRepositoryImpl{
		client: client,
	}
}

func (r *RoleRepositoryImpl) Create(ctx context.Context, input *ent.Role) error {
	_, err := r.client.Role.
		Create().
		SetID(input.ID).
		SetRoleName(input.RoleName).
		SetDescription(input.Description).
		Save(ctx)
	return err
}

func (r *RoleRepositoryImpl) GetList(ctx context.Context) ([]*ent.Role, error) {
	return r.client.Role.
		Query().
		WithPermissions(). // eager load permissions
		WithAccounts().    // eager load accounts
		All(ctx)
}

func (r *RoleRepositoryImpl) Update(ctx context.Context, input *ent.Role) (*ent.Role, error) {
	if input.ID == uuid.Nil {
		return nil, errors.New("invalid role ID")
	}

	return r.client.Role.
		UpdateOneID(input.ID).
		SetRoleName(input.RoleName).
		SetDescription(input.Description).
		Save(ctx)
}

func (r *RoleRepositoryImpl) Delete(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return r.client.Role.
		DeleteOneID(uid).
		Exec(ctx)
}
