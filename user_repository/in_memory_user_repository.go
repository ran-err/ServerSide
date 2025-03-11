package user_repository

import "errors"

type User struct {
	ID     int
	Active bool
	Email  string
}

// InMemoryUserRepository stores a slice of Users by value to:
// 1. Optimize cache efficiency.
// 2. Prevent unintended modifications.
// 3. Handle frequently accessed small structs efficiently.
type InMemoryUserRepository struct {
	Users []User
}

func (r *InMemoryUserRepository) FindUserByEmail(email string) (*User, error) {
	if email == "user@example.com" {
		return &User{Email: "user@example.com"}, nil
	} else {
		return nil, errors.New("user not found")
	}
}
