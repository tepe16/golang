package main

import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

type student struct {
    id    string
    name  string
    age   int
    grade int
}
func connect() (*sql.DB, error) {
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/belajargo")
    if err != nil {
        return nil, err
    }

    return db, nil
}

func sqlQuery() {
    db, err := connect()
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    defer db.Close()

    
    rows, err := db.Query("select id, name, age,grade from tb_student")
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    defer rows.Close()

    var result []student

    for rows.Next() {
        var each = student{}
        var err = rows.Scan(&each.id, &each.name,&each.age, &each.grade)

        if err != nil {
            fmt.Println(err.Error())
            return
        }

        result = append(result, each)
    }

    if err = rows.Err(); err != nil {
        fmt.Println(err.Error())
        return
    }

    for _, each := range result {
		fmt.Println("Nama",each.name)
		fmt.Println("umur",each.age)
		fmt.Println("grade",each.grade)
    }
}
func main() {
    sqlQuery()
}