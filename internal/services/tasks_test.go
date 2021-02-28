package services

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

const (
	MaxTasksInTheDay = 2
)

func TestListTask(t *testing.T) {
	store := mockDB()
	todo := NewToDoService(MaxTasksInTheDay, store)

	handler := http.HandlerFunc(todo.ListTasks)
	req, err := http.NewRequest("GET", "/tasks?created_date=2021-02-28", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assertStatusCode(t, http.StatusOK, rr.Code)
	assertResponseContain(t, "data", rr.Body.String())
}

func TestToDoService(t *testing.T) {
	store := mockDB()
	todo := NewToDoService(MaxTasksInTheDay, store)

	handler := http.HandlerFunc(todo.AddTask)
	body := []byte(`{"content": "Some content"}`)
	req, err := http.NewRequest("POST", "/task", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assertStatusCode(t, http.StatusOK, rr.Code)
	assertResponseContain(t, "Some content", rr.Body.String())
}
