//connecting and disconnecting from the database in question

package datab


import (
	"fmt"
	"log"
	"strconv"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	con "maria-test/config"
)

var DB *sql.DB


func ConnectDB()(*sql.DB, error){
	var err error
	p := con.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32) //(variable, base, size)
	if err != nil {
		log.Println("Int conversion error !!!")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:%d)/%s", con.Config("MYSQL_USER"), con.Config("MYSQL_ROOT_PASSWORD"), port, con.Config("MYSQL_DATABASE"))
    db, err := sql.Open("mysql", dsn)
    if err!= nil {
		fmt.Println(err)
        return nil, err	
    }
    // Ping db to verify the connection
    err = db.Ping()
    if err!= nil {
        fmt.Println(err)
        return nil, err	
    }
    
	//fmt.Println("Connnected to all databases successfuly !!!")
    return db, nil
}