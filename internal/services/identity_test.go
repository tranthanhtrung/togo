package services

import (
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

const (
	JWTKey = "wqGyEBBfPK9w3Lxw"
)

func TestLoginWithValidUser(t *testing.T) {
	store := mockDB()
	identity := NewIdentitySerice(JWTKey, store)
	handler := http.HandlerFunc(identity.GetAuthToken)

	req, err := http.NewRequest("GET", "/login?user_id=firstUser&password=example", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assertStatusCode(t, http.StatusOK, rr.Code)
	assertResponseContain(t, "data", rr.Body.String())
}

func TestLoginWithInalidUser(t *testing.T) {
	store := mockDB()
	identity := NewIdentitySerice(JWTKey, store)
	handler := http.HandlerFunc(identity.GetAuthToken)

	req, err := http.NewRequest("GET", "/login?user_id=firstUs&password=example", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assertStatusCode(t, http.StatusUnauthorized, rr.Code)
	assertResponseContain(t, "incorrect user_id/pwd", rr.Body.String())
}
