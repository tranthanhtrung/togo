package services

import (
	"database/sql"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
)

func TestIdentityService(t *testing.T) {
	db, err := sql.Open("sqlite3", "./../../data.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}

	store := &sqllite.LiteDB{DB: db}
	identity := NewIdentitySerice("wqGyEBBfPK9w3Lxw", store)

	handler := http.HandlerFunc(identity.GetAuthToken)

	// Test Login with valid user
	req, err := http.NewRequest("GET", "/login?user_id=firstUser&password=example", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Login with valid user.\n Expect: %v, actual: %v",
			http.StatusOK, status)
	}
	if !strings.Contains(rr.Body.String(), "data") {
		t.Errorf("Login with valid user.\n Expect: %v, actual: %v", "{data: string}", rr.Body.String())
	}

	// Test Login with invalid user
	req, err = http.NewRequest("GET", "/login?user_id=firstUs&password=example", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("Login with invalid user.\n Expect: %v, actual: %v",
			http.StatusUnauthorized, status)
	}
	if !strings.Contains(rr.Body.String(), "incorrect user_id/pwd") {
		t.Errorf("Login with invalid user.\n Expect: %v, actual: %v", "{data: string}", rr.Body.String())
	}
}
