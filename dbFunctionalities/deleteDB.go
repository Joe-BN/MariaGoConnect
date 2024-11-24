package datab

import (
	"fmt"

	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func DeleteData(id int) error {
	db, _:= ConnectDB()
    // Delete data from the table
    _, err := db.Exec("DELETE FROM Acc WHERE id =?", id)
	db.Close()
	fmt.Println("Deleted user")
    return err
}