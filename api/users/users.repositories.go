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
	err := r.db.Create(user).Error
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

func (r *UserRepository) LoginUser(email *string, username *string, password string) error {
	var user User
    q := r.db.Model(&User{})

    haveCond := false
    if email != nil {
        q = q.Where("email = ?", *email)
        haveCond = true
    }

    if username != nil {
        if haveCond {
            q = q.Or("username = ?", *username)
        } else {
            q = q.Where("username = ?", *username)
            haveCond = true
        }
    }

	var storedHash []byte
	userFound := true

	if err := q.First(&user).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            storedHash = fakeHash
            userFound = false
        } else {
            return err
        }
    } else {
        storedHash = []byte(user.Password)
    }


	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
	if err != nil {
		return errors.New("invalid credentials")
	}

	if (!userFound) {
		return errors.New("invalid credentials")
	}

	return nil
}
