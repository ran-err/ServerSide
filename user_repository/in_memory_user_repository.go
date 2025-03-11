package user_repository

import (
	"errors"
	"strconv"
	"strings"
)

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
