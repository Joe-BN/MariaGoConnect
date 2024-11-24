package main

import (
	"fmt"
	_ "maria-test/config"
	db "maria-test/dbFunctionalities"
)

func main(){
	// Create and connect to the
    fmt.Println("Running test ...")

	/*
	user1 := config.User{
		Name: "Josh",
		Email: "omollo@gmail.com",
		Password: "1234567890",
	}
	db.CreateAcc(user)
	
	db.CreateTable()
	*/
	
	//deleting a user
	db.DeleteData(1)
}
