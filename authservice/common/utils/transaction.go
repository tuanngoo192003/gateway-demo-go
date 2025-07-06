package utils

import (
	"context"

	"github.com/tuanngoo192003/gateway-demo-go/authservice/domain/ent"
	"github.com/tuanngoo192003/gateway-demo-go/authservice/infra/config"
)

func Transaction[T any](ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) (T, error)) (result T, err error) {
	log := config.GetLogger()

	tx, err := client.Tx(ctx)
	if err != nil {
		log.Error("Failed to start transaction: " + err.Error())
		return result, err
	}

	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
			log.Error("Transaction panicked, rolled back")
			panic(r)
		}
	}()

	result, err = fn(tx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Error("Rollback failed: %v", rbErr.Error())
		}
		log.Errorf("Transaction failed, rolled back: %v", err.Error())
		return result, err
	}

	if err := tx.Commit(); err != nil {
		log.Error("Failed to commit transaction: " + err.Error())
		return result, err
	}

	return result, nil
}
