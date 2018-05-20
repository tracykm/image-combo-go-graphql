package main

import (
	"database/sql"
	"log"
)

func insertUser(db *sql.DB, email string) User {
	sqlStatement := `
	INSERT INTO users (email)
	VALUES ($1)
	RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, email).Scan(&id)
	if err != nil {
		panic(err)
	}
	return User{id, email}
}

func queryUsers(db *sql.DB) []*User {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	bks := make([]*User, 0)
	for rows.Next() {
		bk := new(User)
		err := rows.Scan(&bk.ID, &bk.Email)
		if err != nil {
			log.Fatal(err)
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	// for _, bk := range bks {
	// 	fmt.Print(bk.ID, bk.Email)
	// }
	return bks
}
