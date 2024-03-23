package service

import "pwsd_keeper/model"

type Store interface {
	CreatePassword(password model.Password) error
}

type PasswordService struct {
	Pservice Store
}

// CreatePassword Add new password recode in database
func (p PasswordService) CreatePassword(passwordModel model.Password) error {
	err := p.Pservice.CreatePassword(passwordModel)

	return err
}

//
//func GenerateRandomPassword(label string) error {
//
//}
