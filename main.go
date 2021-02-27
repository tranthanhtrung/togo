package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/manabie-com/togo/infra/config"
	"github.com/manabie-com/togo/internal/controller"
	"github.com/manabie-com/togo/internal/services"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
)

func main() {
	cfg, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}
	store := &sqllite.LiteDB{DB: db}

	identity := services.NewIdentitySerice("wqGyEBBfPK9w3Lxw", store)
	todo := services.NewToDoService(cfg.Task.MaxInTheDay, store)

	http.ListenAndServe(":5050", &controller.Handler{
		Identity: identity,
		ToDo:     todo,
	})
}
