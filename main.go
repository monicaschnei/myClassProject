package main

import (
	"database/sql"
	"log"
	"myclass/api"
	db "myclass/db/sqlc"
	"myclass/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config", err)
	}

	connection, err := sql.Open(config.DBDrive, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	store := db.NewStore(connection)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
