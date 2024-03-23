package service

import "pwsd_keeper/model"

type Store interface {
	GetLabel() []string
	CreatePassword(password model.Password) error
}

//
//func CreatePassword(p model.Password) error {
//
//}
//
//func GenerateRandomPassword(label string) error {
//
//}
