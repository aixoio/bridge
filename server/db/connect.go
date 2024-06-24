package db

import (
	"database/sql"
	"fmt"

	"github.com/aixoio/bridge/server/env"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect(dat env.Env) error {
	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dat.Pg_host, dat.Pg_port, dat.Pg_user, dat.Pg_pwd, dat.Pg_db)

	db_conn, err := sql.Open("postgres", url)

	db = db_conn

	return err
}

func Disconnect() {
	db.Close()
}
