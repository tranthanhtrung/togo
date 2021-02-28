package services

import (
	"database/sql"
	"log"
	"strings"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	todosql "github.com/manabie-com/togo/internal/storages/sql"
)

func mockDB() *todosql.Database {
	db, err := sql.Open("sqlite3", "./../../data.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}

	return &todosql.Database{DB: db}
}

func assertStatusCode(t *testing.T, expect, actual int) {
	if expect != actual {
		t.Errorf("Wrong status code. Expect: %v, actual: %v", expect, actual)
	}
}

func assertResponseContain(t *testing.T, expect, actual string) {
	if !strings.Contains(actual, expect) {
		t.Errorf("Response data wrong. Expect constain: %v, actual: %v", expect, actual)
	}
}
