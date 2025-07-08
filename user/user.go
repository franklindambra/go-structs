package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

// receiver: name and type of data with pointer to avoid copy of obj
func (u *User) OutputUserDetails() {
	fmt.Println("Hey", u.firstName, u.lastName, "! Your birthday is", u.birthdate)
}

// the original struct is not changed with these setter operations natively, hence pointer *
func (u *User) ClearUsername() {
	u.firstName = ""
	u.lastName = ""
}

// admin constructor
func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN FIRST NAME",
			lastName:  "ADMIN lAST NAME",
			birthdate: "---",
			createdAt: time.Now(),
		},
	}
}

// constructor function
func New(userFirstName string, userLastName string, userBirthdate string) (*User, error) {

	if userFirstName == "" || userLastName == "" || userBirthdate == "" {
		return nil, errors.New("error")
	}

	return &User{
		userFirstName,
		userLastName,
		userBirthdate,
		time.Now(),
	}, nil
}
