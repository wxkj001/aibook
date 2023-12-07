package app

import (
	"aibook/ent"

	"entgo.io/ent/dialect"
)

func NewDB() *ent.Client {
	client, err := ent.Open(dialect.Postgres, "host=<host> port=<port> user=<user> dbname=<database> password=<pass>")
	if err != nil {
		panic("failed opening connection to postgres")
	}
	return client
}
