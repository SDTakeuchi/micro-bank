package repository

import (
	"main/user/domain/model"
)

type UserRepository interface {
	Get(email string) (model.IUser, error)
}