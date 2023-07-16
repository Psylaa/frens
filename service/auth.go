package service

import "github.com/bwoff11/frens/pkg/database"

type AuthService struct{ Database *database.Database }

func (a *AuthService) Login(username, password string) error {
	return nil
}

func (a *AuthService) Register(username, email, password string) error {
	return nil
}
