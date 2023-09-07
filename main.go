package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/lucasbarretoluz/accountmanagment/api"
	db "github.com/lucasbarretoluz/accountmanagment/db/sqlc"
	"github.com/lucasbarretoluz/accountmanagment/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		panic("Error here")
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		panic(err)
	}
}
