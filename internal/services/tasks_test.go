package services

import (
	"bytes"
	"database/sql"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
)

func TestToDoService(t *testing.T) {
	db, err := sql.Open("sqlite3", "./../../data.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}

	store := &sqllite.LiteDB{DB: db}
	todo := NewToDoService(2, store)

	// Test List task
	handler := http.HandlerFunc(todo.ListTasks)
	req, err := http.NewRequest("GET", "/tasks?created_date=2021-02-27", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Get list tasks.\n Expect: %v, actual: %v",
			http.StatusOK, status)
	}
	if !strings.Contains(rr.Body.String(), "data") {
		t.Errorf("Get list tasks.\n Expect: %v, actual: %v", "{data: Array}", rr.Body.String())
	}

	// Test Add task
	handler = http.HandlerFunc(todo.AddTask)
	body := []byte(`{"content": "Some content"}`)
	req, err = http.NewRequest("POST", "/task", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Add task.\n Expect: %v, actual: %v",
			http.StatusOK, status)
	}
	if !strings.Contains(rr.Body.String(), "Some content") {
		t.Errorf("Add task.\n Expect: %v, actual: %v", "{data: string}", rr.Body.String())
	}
}
