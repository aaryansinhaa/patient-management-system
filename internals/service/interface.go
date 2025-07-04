package service

import "github.com/aaryansinhaa/patient-management-system/internals/model"

type AuthService interface {
	Register(user *model.User) error
	Login(username, password string) (*model.User, string, error)
}
