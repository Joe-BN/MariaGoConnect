package datab


import (
    "fmt"
)

func SearchAcc(iden string) error {
	db,_ := ConnectDB()
	// Prepare the SQL query
    stmt, err := db.Prepare("SELECT * FROM Acc WHERE id =?")
    if err!= nil {
        fmt.Println(err)
        return err
    }
    defer stmt.Close()

    // Execute the query with the ID parameter
    rows, err := stmt.Query(iden)
    if err!= nil {
        fmt.Println(err)
        return err
    }
    defer rows.Close()

    // Fetch the result
    for rows.Next() {
        var (
			id int
            name   string
            email string
            password    string
        )
        err := rows.Scan(&id, &name, &email, &password)
        if err!= nil {
            fmt.Println(err)
            return err
        }
        fmt.Printf("ID: %d, Name: %s, Email: %s Password %s\n", id, name, email, password)
    }
	return err
}