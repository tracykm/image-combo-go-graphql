package main

import (
	"database/sql"
	"log"

	"github.com/graphql-go/graphql"
)

type User struct {
	ID    int    `json:"id"`
	Email string `json:"name"`
}

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
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
