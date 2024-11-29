package datab

import (
    "fmt"
    "log"

	"database/sql"
    _ "github.com/go-sql-driver/mysql"

	con "maria-test/config"
)

var DB *sql.DB

func ConnectDB() (*sql.DB, error) {
    dsn := fmt.Sprintf("%s:%s@tcp(localhost:%s)/%s",
        con.Config("MYSQL_USER"),
        con.Config("MYSQL_ROOT_PASSWORD"),
        con.Config("DB_PORT"),
        con.Config("MYSQL_DATABASE"),
    )

    db, err := sql.Open("mysql", dsn)
    if err!= nil {
        log.Printf("Error connecting to database: %v", err)
        return nil, err
    }

    err = db.Ping()
    if err!= nil {
        log.Printf("Error pinging database: %v", err)
        return nil, err
    }

    //log.Println("Connected to database successfully")
    return db, nil
}