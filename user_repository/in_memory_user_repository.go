package user_repository

import "errors"

type User struct {
	Email string
}

type InMemoryUserRepository struct {
}

func (r *InMemoryUserRepository) FindUserByEmail(email string) (*User, error) {
	if email == "user@example.com" {
		return &User{Email: "user@example.com"}, nil
	} else {
		return nil, errors.New("user not found")
	}
}
