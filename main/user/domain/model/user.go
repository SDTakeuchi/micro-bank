package model

import (
	"crypto/sha256"
	"fmt"
	"time"

	"google/uuid"

	"github.com/go-playground/validator/v10"
)

var (
	minPasswordLength = 8
)

type user struct {
	id             string `validate:"required"`
	email          string `validate:"required,email"`
	hashedPassword string `validate:"required"`
	firstName      string `validate:"required"`
	familyName     string `validate:"required"`
	kycDone        bool
	createdAt      time.Time `validate:"required"`
	updatedAt      *time.Time
	deletedAt      *time.Time
}

type IUser interface {
	ID() string
	Email() string
	HashedPassword() string
	FirstName() string
	FamilyName() string
	KYCDone() bool
	CreatedAt() time.Time
	UpdatedAt() time.Time
	DeletedAt() time.Time

	FullName() string
	ProceedKYC()
}

func (u *user) ID() string             { return u.id }
func (u *user) Email() string          { return u.email }
func (u *user) HashedPassword() string { return u.hashedPassword }
func (u *user) FirstName() string      { return u.firstName }
func (u *user) FamilyName() string     { return u.familyName }
func (u *user) KYCDone() bool          { return u.kycDone }
func (u *user) CreatedAt() time.Time   { return u.createdAt }
func (u *user) UpdatedAt() time.Time   { return *u.updatedAt }
func (u *user) DeletedAt() time.Time   { return *u.deletedAt }

func (u *user) FullName() string {
	return fmt.Sprintf("%s %s", u.firstName, u.familyName)
}

func (u *user) ProceedKYC() {
	u.kycDone = true
}

func NewUser(firstName, familyName, email, rawPassword string) (IUser, error) {
	if len(rawPassword) < minPasswordLength {
		return nil, fmt.Errorf(
			"failed to create user: length of password must be longer than %d",
			minPasswordLength,
		)
	}
	hashedPassword := sha256.Sum256([]byte(rawPassword))
	u := &user{
		id:             uuid.NewV4(),
		email:          email,
		firstName:      firstName,
		familyName:     familyName,
		hashedPassword: string(hashedPassword[:]),
		kycDone:        false,
		createdAt:      time.Now(),
		updatedAt:      nil,
		deletedAt:      nil,
	}
	validate := validator.New()
	if errors := validate.Struct(u); errors != nil {
		return nil, errors
	}
	return u, nil
}
