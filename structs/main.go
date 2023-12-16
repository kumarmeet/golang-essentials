package main

import (
	"fmt"

	"example.com/structs/custom_types"
	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	//now here New make it clear we are creating new user
	//appUser inferred by pointer variable automatically by this short hand syntax :=
	appUser, error := user.New(userFirstName, userLastName, userBirthdate)

	if error != nil {
		panic(error)
	}

	appUser.OutputUserDetails()

	appUser.ClearUserName()

	appUser.OutputUserDetails()

	email := getUserData("Please enter your email: ")
	password := getUserData("Please enter your password: ")

	//appAdmin inferred by pointer variable automatically by this short hand syntax :=
	appAdmin, error := user.NewAdmin(email, password, appUser)

	if error != nil {
		panic(error)
	}

	appAdmin.OutputUserDetails()

	var name custom_types.Str = "Meet Kumar Vishwakarma"
	name.Log()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
