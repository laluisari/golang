package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/model"
)

func isExpired(s model.Session) bool {
	return s.Expiry.Before(time.Now())
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("session_token")
		if err != nil {
			w.WriteHeader(401)
			jsonData, _ := json.Marshal(model.ErrorResponse{
				Error: "http: named cookie not present",
			})
			w.Write(jsonData)
			return
		}
		stoken := c.Value
		userSession, ok := db.Sessions[stoken]

		if !ok {
			w.WriteHeader(401)
			jsonData, _ := json.Marshal(model.ErrorResponse{
				Error: "http: named cookie not present",
			})
			w.Write(jsonData)
			return
		}

		if isExpired(userSession) {
			delete(db.Sessions, stoken)
			w.WriteHeader(401)
			jsonData, _ := json.Marshal(model.ErrorResponse{
				Error: "http: named cookie not present",
			})
			w.Write(jsonData)
			return
		}

		ctx := context.WithValue(r.Context(), "username", userSession.Username)
		next.ServeHTTP(w, r.WithContext(ctx))

	}) // TODO: replace this
}
