// package main

// import (
// "database/sql"
// "fmt"

// \_ "github.com/lib/pq"
// )

// const (
// host = "localhost"
// port = 5432
// user = "tracy"
// password = "boo"
// dbname = "image_combos"
// )

// func main() {
// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// "password=%s dbname=%s sslmode=disable",
// host, port, user, password, dbname)
// db, err := sql.Open("postgres", psqlInfo)
// if err != nil {
// panic(err)
// }
// defer db.Close()

// err = db.Ping()
// if err != nil {
// panic(err)
// }
// sqlStatement := `// INSERT INTO users (email) // VALUES ($1) // RETURNING id`
// id := 0
// err = db.QueryRow(sqlStatement, "t@m.com").Scan(&id)
// if err != nil {
// panic(err)
// }
// fmt.Println("New record ID is:", id)
// }
