package entity

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id             uint
	Username       string
	Role           string
	HashedPassword string
	CreatedAt      time.Time
}

func NewUser(username, password, role string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("cannot hash password: %w", err)
	}

	user := &User{
		Username:       username,
		Role:           role,
		HashedPassword: string(hashedPassword),
		CreatedAt:      time.Now(),
	}

	return user, nil
}

func (user *User) IsPasswordCorrect(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	return err == nil
}
