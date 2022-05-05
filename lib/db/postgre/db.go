package posgre

import (
	"chatroom/conf"
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func ConDb() (db *sql.DB, err error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.DB_HOST, conf.DB_PORT, conf.DB_USER, conf.DB_PASSWORD, conf.DB_NAME)
	db, err = sql.Open(conf.DB_DRIVER, connStr)
	if err != nil {
		return
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		return
	}
	return
}
