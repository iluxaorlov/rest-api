package store

import "github.com/iluxaorlov/rest-api/internal/app/model"

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(string) (*model.User, error)
}