/*
@Time : 2022/1/23 18:57
@Author : Hwdhy
@File : user
@Software: GoLand
*/
package service

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username       string
	HashedPassword string
	Role           string
}

func NewUser(username string, password string, role string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("cannot hash password: %w", err)
	}

	user := &User{
		username,
		string(hashedPassword),
		role,
	}
	return user, nil
}

func (user *User) IsCorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	return err == nil
}

func (user *User) Clone() *User {
	return &User{
		Username:       user.Username,
		HashedPassword: user.HashedPassword,
		Role:           user.Role,
	}
}