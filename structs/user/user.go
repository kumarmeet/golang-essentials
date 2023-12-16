package user

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// struct embedding
type admin struct {
	email    string
	password string
	user     //anonymous struct embedding
	// User user -> named struct embedding
}

// method of user struct, and we dont need to pass address, due to this method is belongs directly to user struct
// func (u user) special reciever argument not a normal agruments
func (u *user) outputUserFullName() string {
	return fmt.Sprintf("%s %s", u.firstName, u.lastName)
}

// mutation methods
func (u *user) ClearUserName() {
	// we cant omit/dereferencing (*u).FirstName and so on, this conversion will done by golang
	u.firstName = ""
	u.lastName = ""
}

// creation or constructor function
// prefix new is a convention to simplify that this is a constructor or creation function (as also utility)
// func NewUser(firstName, lastName, birthdate string) (*user, error) {
// 	if firstName == "" || lastName == "" || birthdate == "" {
// 		return nil, errors.New("First name, last name and birth date is required!")
// 	}

//		return &user{
//			firstName: firstName,
//			lastName:  lastName,
//			birthdate: birthdate,
//			createdAt: time.Now(),
//		}, nil
//	}

// when seprate the struct and all belonging methods into package then we can use New
func New(firstName, lastName, birthdate string) (*user, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First name, last name and birth date is required!")
	}

	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string, u *user) (*admin, error) {
	if email == "" || password == "" {
		return nil, errors.New("Email and password required!")
	}

	return &admin{
		email:    email,
		password: password,
		user: user{
			firstName: u.firstName,
			lastName:  u.lastName,
			birthdate: u.birthdate,
			createdAt: u.createdAt,
		},
	}, nil
}

func (u user) OutputUserDetails() {
	fmt.Println(u.outputUserFullName(), u.firstName, u.createdAt)
}

func (a admin) OutputUserDetails() {
	fmt.Println(a.email, a.password)
}

//this means instantiate the user struct with nil value
// appUser := user{}

//omit the key, and just specify the values. but make sure order is the same as in the blueprint as struct
// appUser := user{
// 	userFirstName,
// 	userLastName,
// 	userBirthdate,
// 	time.Now(),
// }

// appUser := user{
// 	firstName: userFirstName,
// 	lastName:  userLastName,
// 	birthdate: userBirthdate,
// 	createdAt: time.Now(),
// }
