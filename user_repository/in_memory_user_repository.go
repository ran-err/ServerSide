package user_repository

import (
	"errors"
	"strconv"
	"strings"
)

type UserID struct {
	ID int
}

type UserStatus struct {
	Active bool
}

type UserEmail struct {
	Email string
}

type User struct {
	UserID
	UserStatus
	UserEmail
}

func NewUser(id int, active bool, email string) User {
	return User{UserID: UserID{ID: id}, UserStatus: UserStatus{Active: active}, UserEmail: UserEmail{Email: email}}
}

// InMemoryUserRepository stores a slice of Users by value to:
// 1. Optimize cache efficiency.
// 2. Prevent unintended modifications.
// 3. Handle frequently accessed small structs efficiently.
type InMemoryUserRepository struct {
	Users []User
}

func New() *InMemoryUserRepository {
	return &InMemoryUserRepository{Users: []User{}}
}

func NewFromSlice(users []User) *InMemoryUserRepository {
	return &InMemoryUserRepository{Users: users}
}

func (r *InMemoryUserRepository) FindUserByEmail(email string) (*User, error) {
	for _, user := range r.Users {
		if user.Email == email && user.Active {
			result := NewUser(user.ID, user.Active, user.Email)
			return &result, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *InMemoryUserRepository) Len() int {
	return len(r.Users)
}

func (r *InMemoryUserRepository) Peek() string {
	var output strings.Builder
	for _, user := range r.Users {
		output.WriteString("{")
		if user.Active {
			output.WriteString("+")
		} else {
			output.WriteString("-")
		}
		output.WriteString(strconv.Itoa(user.ID))
		output.WriteString(", ")
		output.WriteString(user.Email)
		output.WriteString("}\n")
	}
	if output.Len() > 0 {
		return output.String()[:output.Len()-1]
	}
	return output.String()
}
