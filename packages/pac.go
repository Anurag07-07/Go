package main

//Syntax
// go mod init github.com/golang


//to install package
//go get package_name
//go mod tidy
import (
	"github.com/golang/auth"
	"github.com/golang/user"
)

func main() {
	auth.LoginWithCredentials("Anurag","772002")

	user:=user.User{
		Email: "Anurag@gmail.com",
		Name:"Anurag",
	}

	println(user.Email)
}