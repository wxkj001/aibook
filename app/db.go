package app

import (
	"aibook/ent"
	"context"
	"log"
	"strings"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func NewDB() *ent.Client {
	client, err := ent.Open(dialect.Postgres, strings.Join([]string{
		"host=" + viper.GetString("db.host"),
		"port=" + viper.GetString("db.port"),
		"user=" + viper.GetString("db.user"),
		"dbname=" + viper.GetString("db.dbname"),
		"password=" + viper.GetString("db.password"),
		"sslmode=disable",
	}, " "))
	if err != nil {
		panic("failed opening connection to postgres：" + err.Error())
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
