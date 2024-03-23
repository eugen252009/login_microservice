package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var db *pgx.Conn
var Ctx context.Context

func Connect(ctx *context.Context) {
	if ctx == nil {
		Ctx = context.TODO()
	}
	dburl := os.Getenv("DB_URL")
	conn, err := pgx.Connect(Ctx, dburl)
	if err != nil {
		log.Fatalf(err.Error())
	}
	db = conn
	err = db.Ping(Ctx)
	if err != nil {
		log.Fatalf("error Pinging db")
	}
}
func GetUser(id string) pgx.Row {
	return db.QueryRow(Ctx, "select id,first_name,last_name from nutzer where id=$1", id)
}
