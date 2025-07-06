package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/tuanngoo192003/gateway-demo-go/authservice/common/utils"
	"github.com/tuanngoo192003/gateway-demo-go/authservice/domain/ent"
	"github.com/tuanngoo192003/gateway-demo-go/authservice/domain/ent/account"
)

type AccountRepository interface {
	Create(ctx context.Context, input *ent.Account) error
	GetList(ctx context.Context) ([]*ent.Account, error)
	Update(ctx context.Context, input *ent.Account) (*ent.Account, error)
	Delete(ctx context.Context, id string) error
	FindOneByFields(ctx context.Context, id string) (*ent.Account, error)
	FindManyByFields(ctx context.Context, fields map[string]any) ([]*ent.Account, error)
}

type AccountRepositoryImpl struct {
	client *ent.Client
}

func NewAccountRepository(client *ent.Client) *AccountRepositoryImpl {
	return &AccountRepositoryImpl{client: client}
}


func (a *AccountRepositoryImpl) Create(ctx context.Context, input *ent.Account) error {
	_, err := utils.Transaction(ctx, a.client, func(tx *ent.Tx) (*ent.Account, error) {
		roleID, err := input.QueryRole().OnlyID(ctx)
		if err != nil {
			return nil, err
		}

		return tx.Account.
			Create().
			SetID(input.ID).
			SetUsername(input.Username).
			SetPassword(input.Password).
			SetEmail(input.Email).
			SetPhoneNumber(input.PhoneNumber).
			SetOAuthType(input.OAuthType).
			SetRoleID(roleID).
			Save(ctx)
	})

	return err
}


func (a *AccountRepositoryImpl) GetList(ctx context.Context) ([]*ent.Account, error) {
	return a.client.Account.
		Query().
		WithRole().
		All(ctx)
}


func (a *AccountRepositoryImpl) Update(ctx context.Context, input *ent.Account) (*ent.Account, error) {
	if input.ID == uuid.Nil {
		return nil, errors.New("invalid account ID")
	}

	return utils.Transaction(ctx, a.client, func(tx *ent.Tx) (*ent.Account, error) {
		roleID, err := input.QueryRole().OnlyID(ctx)
		if err != nil {
			return nil, err
		}

		return tx.Account.
			UpdateOneID(input.ID).
			SetUsername(input.Username).
			SetPassword(input.Password).
			SetEmail(input.Email).
			SetPhoneNumber(input.PhoneNumber).
			SetOAuthType(input.OAuthType).
			SetRoleID(roleID).
			Save(ctx)
	})
}


func (a *AccountRepositoryImpl) Delete(ctx context.Context, id string) error {
	_, err := utils.Transaction(ctx, a.client, func(tx *ent.Tx) (any, error) {
		uid, err := uuid.Parse(id)
		if err != nil {
			return nil, err
		}

		return nil, tx.Account.
			DeleteOneID(uid).
			Exec(ctx)
	})

	return err
}

func (a *AccountRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*ent.Account, error) {		
	return a.client.Account.Query().Where(account.IDEQ(id)).Only(ctx)
}

func (a *AccountRepositoryImpl)	FindManyByFields(ctx context.Context, fields map[string]any) ([]*ent.Account, error) {
	query := a.client.Account.Query()

	for field, value := range fields {
		switch field {
		case "id":
			if id, ok := value.(uuid.UUID); ok {
				query = query.Where(account.IDEQ(id))
			}
		case "username":
			if username, ok := value.(string); ok {
				query = query.Where(account.UsernameEQ(username))
			}
		case "email":
			if email, ok := value.(string); ok {
				query = query.Where(account.EmailEQ(email))
			}
		case "phoneNumber":
			if phone, ok := value.(string); ok {
				query = query.Where(account.PhoneNumberEQ(phone))
			}
		case "oAuthType":
			if oAuth, ok := value.(account.OAuthType); ok {
				query = query.Where(account.OAuthTypeEQ(oAuth))
			}
		default:
			return nil, &ent.NotFoundError{}	
		}
	}

	return query.All(ctx)
}







