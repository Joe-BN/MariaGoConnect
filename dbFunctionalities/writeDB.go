package datab

import (
	"maria-test/config"
)

func CreateAcc(user config.User) error{
	db,_:= ConnectDB()
	query := "INSERT INTO Acc (name, email, password) VALUES (?,?,?)"
	_, err := db.Exec(query, user.Name, user.Email, user.Password)
	db.Close()
	return err
}

