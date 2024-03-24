package service

import (
	"pwsd_keeper/model"
)

type PasswordStore interface {
	CreatePassword(password *model.Password) error
}

type PasswordService struct {
	Repo PasswordStore
}

// CreatePassword Add new password recode in database
func (p PasswordService) CreatePassword(passwordModel *model.Password) error {
	err := p.Repo.CreatePassword(passwordModel)

	return err
}
