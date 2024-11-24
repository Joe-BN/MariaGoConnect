package datab

import(
	"fmt"

	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func CreateTable()error{
	db, _:= ConnectDB()
	_ ,err := db.Exec(`
		CREATE TABLE IF NOT EXISTS Acc (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255),
			email VARCHAR(255),
			password VARCHAR(255)
		);
	`)
	fmt.Println("User database created !")
	db.Close()
	return err
}