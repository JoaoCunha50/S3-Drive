package users

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

var fakeHash []byte

func init() {
	fakeHash, _ = bcrypt.GenerateFromPassword([]byte("fakers mans"), bcrypt.DefaultCost)
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *User) error {
	err := r.db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetUser(id int) (*User, error) {
	var user User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) LoginUser(email string, username string, password string) error {
	var user User
	var storedHash []byte
	userFound := true

	err := r.db.First(&user, "email = ? OR username = ?", email, username).Error

	switch (err) {
	case gorm.ErrRecordNotFound:
		storedHash = fakeHash
		userFound = false
	case nil:
		storedHash = []byte(user.Password)
	default:
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
	if err != nil {
		return errors.New("invalid credentials")
	}

	if (!userFound) {
		return errors.New("invalid credentials")
	}

	return nil
}
