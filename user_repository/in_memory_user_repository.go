package user_repository

import "errors"

type InMemoryUserRepository struct {
}

func (r *InMemoryUserRepository) FindUserByEmail(email string) (*int, error) {
	return nil, errors.New("user not found")
}
