package utils

import (
	"context"
	"database/sql"
	"net/http"
)

// UserAuthKey type user auth key
type UserAuthKey int8

// UserIDFromCtx get UserID from context
func UserIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(UserAuthKey(0))
	id, ok := v.(string)
	return id, ok
}

// ValueFromHTTPRequest get value with key from request
func ValueFromHTTPRequest(req *http.Request, p string) sql.NullString {
	return sql.NullString{
		String: req.FormValue(p),
		Valid:  true,
	}
}
