package datab

import (
    "fmt"
    "log"
)

// SearchAcc retrieves account data by ID
func SearchAcc(iden string) error {
    db, err := ConnectDB()
    if err!= nil {
        return err
    }
    defer db.Close()

    stmt, err := db.Prepare("SELECT * FROM Acc WHERE id =?")
    if err!= nil {
        log.Println(err)
        return err
    }
    defer stmt.Close()

    rows, err := stmt.Query(iden)
    if err!= nil {
        log.Println(err)
        return err
    }
    defer rows.Close()

    var (
        id   int
        name string
        email string
        password string
    )
    var errors []error

    for rows.Next() {
        err = rows.Scan(&id, &name, &email, &password)
        if err!= nil {
            errors = append(errors, err)
            continue
        }
        fmt.Printf("ID: %d, Name: %s, Email: %s Password %s\n", id, name, email, password)
    }

    if len(errors) > 0 {
        return errors[0]
    }
    return nil
}