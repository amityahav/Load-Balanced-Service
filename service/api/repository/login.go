package repository

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"loadBalancedService/ent"
	"loadBalancedService/ent/user"
	"net/http"
	"time"
)

type LoginService struct {
	db *ent.Client
}

func (ls *LoginService) Login(ctx context.Context, username string, password string) (*http.Cookie, error) {
	usr, err := ls.db.User.
		Query().
		Where(user.Username(username)).
		Only(ctx)

	if err != nil {
		return &http.Cookie{}, errors.Wrap(err, "failed querying user")
	}

	if DoPasswordsMatch(password, usr.Salt, usr.Password) {
		newCookie := http.Cookie{
			Name:    "username",
			Value:   username,
			MaxAge:  86400,
			Expires: time.Now().Add(30 * time.Minute),
		}

		return &newCookie, nil
	}

	return &http.Cookie{}, fmt.Errorf("invalid password")

}
