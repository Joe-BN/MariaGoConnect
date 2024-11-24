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

func UpdateAcc(id int, user config.User) error {
    db, _:= ConnectDB()
    _, err := db.Exec(`
        UPDATE Acc
        SET name =?, email =?, password =?
        WHERE id =?;
    `, user.Name, user.Email, user.Password, id)
    return err
}