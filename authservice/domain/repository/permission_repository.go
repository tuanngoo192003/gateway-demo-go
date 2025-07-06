package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/tuanngoo192003/gateway-demo-go/authservice/domain/ent"
)

type PermissionRepository interface {
	Create(ctx context.Context, input *ent.Permission) error
	GetList(ctx context.Context) ([]*ent.Permission, error)
	Update(ctx context.Context, input *ent.Permission) (*ent.Permission, error)
	Delete(ctx context.Context, id string) error
}

type PermissionRepositoryImpl struct {
	client *ent.Client
}

func NewPermissionRepository(client *ent.Client) *PermissionRepositoryImpl {
	return &PermissionRepositoryImpl{
		client: client,
	}
}

func (r *PermissionRepositoryImpl) Create(ctx context.Context, input *ent.Permission) error {
	_, err := r.client.Permission.
		Create().
		SetID(input.ID).
		SetPermissionName(input.PermissionName).
		SetURI(input.URI).
		SetMethod(input.Method).
		Save(ctx)
	return err
}

func (r *PermissionRepositoryImpl) GetList(ctx context.Context) ([]*ent.Permission, error) {
	return r.client.Permission.
		Query().
		All(ctx)
}

func (r *PermissionRepositoryImpl) Update(ctx context.Context, input *ent.Permission) (*ent.Permission, error) {
	if input.ID == uuid.Nil {
		return nil, errors.New("invalid permission ID")
	}

	return r.client.Permission.
		UpdateOneID(input.ID).
		SetPermissionName(input.PermissionName).
		SetURI(input.URI).
		SetMethod(input.Method).
		Save(ctx)
}

func (r *PermissionRepositoryImpl) Delete(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return r.client.Permission.
		DeleteOneID(uid).
		Exec(ctx)
}
