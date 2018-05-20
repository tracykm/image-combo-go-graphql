package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	_ "github.com/lib/pq"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/mnmtanish/go-graphiql"
	"github.com/rs/cors"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "tracy"
	password = "boo"
	dbname   = "image_combos"
)

type User struct {
	ID    int    `json:"id"`
	Email string `json:"name"`
}

func init() {
	rand.Seed(time.Now().UnixNano())

	db := initDb()
	defer db.Close()
	// insertUser(db, "e@b.com")
	users := queryUsers(db)
	println(users[0].Email)
}

func initDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
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

// // var Users []User
// func getUsers() []User {
// 	results := UserList
// 	return results
// }

var personType = graphql.NewObject(
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
var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"Users": &graphql.Field{
			Type:        graphql.NewList(personType),
			Description: "List of people",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return make([]*User, 0), nil
			},
		},
	},
})

// root mutation
var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		/*
			curl -g 'http://localhost:8080/graphql?query=mutation+_{createUser(text:"My+new+person"){id,text,done}}'
		*/
		"createUser": &graphql.Field{
			Type:        personType, // the return type for this field
			Description: "Create new person",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				// marshall and cast the argument value
				name, _ := params.Args["name"].(string)

				// perform mutation operation here
				// for e.g. create a User and save to DB.
				newUser := User{
					ID:    4,
					Email: name,
				}

				// UserList = append(UserList, newUser)

				// return the new User object that we supposedly save to DB
				// Note here that
				// - we are returning a `User` struct instance here
				// - we previously specified the return Type to be `personType`
				// - `User` struct maps to `personType`, as defined in `personType` ObjectConfig`
				return newUser, nil
			},
		},
	},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    queryType,
	Mutation: rootMutation,
})

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})
	h := handler.New(&handler.Config{
		Schema: &Schema,
		Pretty: true,
	})
	// serve HTTP
	serveMux := http.NewServeMux()
	// serveMux.HandleFunc("/neo", neo4jHandler)
	serveMux.Handle("/graphql", c.Handler(h))
	serveMux.HandleFunc("/graphiql", graphiql.ServeGraphiQL)
	http.ListenAndServe(":8080", serveMux)
}

// query {
// 	Users {
// 	  name
// 	  id
// 	}
//   }

// mutation {
// 	createUser(name: "Lenny") {
// 	  name
// 	  id
// 	}
//   }
