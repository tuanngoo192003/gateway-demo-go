package database

import (
	"database/sql"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/tuanngoo192003/gateway-demo-go/authservice/domain/ent"
	"github.com/tuanngoo192003/gateway-demo-go/authservice/infra/config"
)

type Database struct {
	Client *ent.Client
}

func NewDatabase(connectionString string) (*Database, error) {
	//Open raw database connection
	log := config.GetLogger()
	log.Info(connectionString)

	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		log.Fatal(err)

		return nil, err
	}

	driver := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(driver))

	return &Database{Client: client}, nil
}
