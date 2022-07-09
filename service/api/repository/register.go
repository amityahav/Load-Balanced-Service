package repository

import (
	"context"
	"github.com/pkg/errors"
	"loadBalancedService/ent"
	"log"
)

type RegisterService struct {
	db *ent.Client
}

func (rs *RegisterService) Register(ctx context.Context, username string, password string) error {
	salt, err := GenerateRandomSalt()
	if err != nil {
		return errors.Wrap(err, "failed creating a new user")
	}

	hashedPassword := HashPassword(password, salt)
	user, err := rs.db.User.
		Create().
		SetUsername(username).
		SetPassword(hashedPassword).
		SetSalt(salt).
		Save(ctx)

	if err != nil {
		return errors.Wrap(err, "failed creating a new user")
	}

	log.Println("user created: ", user)
	return nil
}
