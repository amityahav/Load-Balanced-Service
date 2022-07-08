package repository

import "loadBalancedService/ent"

type RegisterService struct {
	db *ent.Client
}

func (rs *RegisterService) Register(username string, password string) error {
	//TODO will handle the DB part of the registration
	return nil
}
