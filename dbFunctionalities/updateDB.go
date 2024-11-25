package datab

import (
	"maria-test/config"
)

func UpdateAcc(id int, user config.User) error {
    db, _:= ConnectDB()
    _, err := db.Exec(`
        UPDATE Acc
        SET name =?, email =?, password =?
        WHERE id =?;
    `, user.Name, user.Email, user.Password, id)
    return err
}