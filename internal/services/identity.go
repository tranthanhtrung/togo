package services

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	todoerr "github.com/manabie-com/togo/internal/httperror"
	"github.com/manabie-com/togo/internal/storages/sql"
	"github.com/manabie-com/togo/utils"
)

// IdentityService implement HTTP server
type IdentityService struct {
	JWTKey string
	Store  *sql.Database
}

// NewIdentitySerice create new IdentityService
func NewIdentitySerice(jwtKey string, store *sql.Database) *IdentityService {
	return &IdentityService{
		JWTKey: jwtKey,
		Store:  store,
	}
}

// GetAuthToken get token of a user
func (s *IdentityService) GetAuthToken(resp http.ResponseWriter, req *http.Request) {
	id := utils.ValueFromHTTPRequest(req, "user_id")
	pass := utils.ValueFromHTTPRequest(req, "password")
	if !s.Store.ValidateUser(req.Context(), id, pass) {
		todoerr.WrapError(resp, http.StatusUnauthorized, "incorrect user_id/pwd")
		return
	}

	token, err := s.createToken(id.String)
	if err != nil {
		todoerr.WrapError(resp, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(resp).Encode(map[string]string{
		"data": token,
	})
}

// ValidToken validate token of a user
func (s *IdentityService) ValidToken(req *http.Request) (*http.Request, bool) {
	token := req.Header.Get("Authorization")

	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(s.JWTKey), nil
	})
	if err != nil {
		log.Println(err)
		return req, false
	}

	if !t.Valid {
		return req, false
	}

	id, ok := claims["user_id"].(string)
	if !ok {
		return req, false
	}

	req = req.WithContext(context.WithValue(req.Context(), utils.UserAuthKey(0), id))
	return req, true
}

func (s *IdentityService) createToken(id string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = id
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(s.JWTKey))
	if err != nil {
		return "", err
	}
	return token, nil
}
