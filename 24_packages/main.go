package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/satyam-code45/go/auth"
	"github.com/satyam-code45/go/user"
)

func main() {
	auth.LoginWithCredentials("satyam", "1234")
	session := auth.GetSession()
	fmt.Println(session)

	user := user.User{
		Email: "satyam@gmail.com",
		Name: "Satyam",
	}
	// fmt.Println(user.Email, user.Name)

	color.Red(user.Email)
	color.Green(user.Name)
}
