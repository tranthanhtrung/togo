package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"github.com/manabie-com/togo/infra/config"
	"github.com/manabie-com/togo/infra/database"
	"github.com/manabie-com/togo/internal/controller"
	"github.com/manabie-com/togo/internal/services"
	todosql "github.com/manabie-com/togo/internal/storages/sql"
)

func main() {
	cfg, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	db, err := database.NewDatabase(cfg)
	if err != nil {
		log.Fatal("cannot opening db:", err)
	}
	store := &todosql.Database{DB: db}

	identity := services.NewIdentitySerice("wqGyEBBfPK9w3Lxw", store)
	todo := services.NewToDoService(cfg.Task.MaxInTheDay, store)

	http.ListenAndServe(":5050", &controller.Handler{
		Identity: identity,
		ToDo:     todo,
	})
}
