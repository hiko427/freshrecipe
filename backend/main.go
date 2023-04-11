package main

import (
	"database/sql"
	"log"

	"github.com/hiko427/myrecipe/api"
	db "github.com/hiko427/myrecipe/db/sqlc"
	util "github.com/hiko427/myrecipe/utils"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server := api.Newserver(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
