package handlers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"time"
)

type authHandler struct {
	secret []byte
	apiKey string
}

func NewAuthHandler() *authHandler {
	return &authHandler{
		secret: []byte("SUPER SECRET"),
		apiKey: "1234",
	}
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header["Api_key"] != nil {
		if r.Header["Api_key"][0] == h.apiKey {
			token, err := h.createJWT()
			if err != nil {
				log.Println(err.Error())
				w.WriteHeader(http.StatusBadRequest)
				_, _ = w.Write([]byte(err.Error()))
				return
			}
			_, _ = w.Write([]byte(fmt.Sprintf("New JWT: %s", token)))
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Invalid api key"))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("api key must be provided"))
	}

}

func (h *authHandler) createJWT() (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	tokenStr, err := token.SignedString(h.secret)
	if err != nil {
		return "", errors.Wrap(err, "failed creating a new user")
	}

	return tokenStr, nil
}

func (h *authHandler) ValidateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					w.WriteHeader(http.StatusUnauthorized)
					_, _ = w.Write([]byte("Not Authorized"))
				}
				return h.secret, nil
			})

			if err != nil || !token.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				_, _ = w.Write([]byte("Not Authorized"))
				return
			}

			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("Token is required"))
		}
	})
}
