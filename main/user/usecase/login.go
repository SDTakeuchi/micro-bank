package usecase

import (
	"fmt"
	"main/user/domain/model"
	"main/user/domain/repository"

	"gorm.io/gorm"
)

type (
	Login interface {
		Do(in LoginInput) (*LoginOutput, error)
	}
	login struct {
		db       *gorm.DB
		userRepo repository.UserRepository
	}
	LoginInput struct {
		Email    string
		Password string
	}
	LoginOutput struct {
		Token string
	}
)

func NewLogin(
	db *gorm.DB,
	userRepo repository.UserRepository,
) (Login) {
	return &login{db: db, userRepo: userRepo}
}

func (u *login) Do(in LoginInput) (*LoginOutput, error) {
	fetchedUser, err := u.userRepo.Get(in.Email)
	if err != nil {
		return nil, err
	}
	err = model.ComparePassword(in.Password, fetchedUser.HashedPassword())
	if err != nil {
		return nil, fmt.Errorf("email and/or password is incorrect")
	}
	return 
}
