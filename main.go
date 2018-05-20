package main

import (
	"database/sql"
	"fmt"
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

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"Users": &graphql.Field{
			Type:        graphql.NewList(userType),
			Description: "List of people",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				db := initDb()
				defer db.Close()
				users := queryUsers(db)
				return users, nil
			},
		},
		"SrcImages": &graphql.Field{
			Type:        graphql.NewList(srcType),
			Description: "List of images",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				db := initDb()
				defer db.Close()
				images := querySrcImages(db)
				return images, nil
			},
		},
	},
})

// root mutation
var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		/*
			curl -g 'http://localhost:8080/graphql?query=mutation+_{createUser(text:"My+new+user"){id,text,done}}'
		*/
		"createUser": &graphql.Field{
			Type:        userType, // the return type for this field
			Description: "Create new user",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				// marshall and cast the argument value
				name, _ := params.Args["name"].(string)

				db := initDb()
				defer db.Close()
				newUser := insertUser(db, name)
				return newUser, nil
			},
		},
		"createSrcImage": &graphql.Field{
			Type:        srcType, // the return type for this field
			Description: "Create new image",
			Args: graphql.FieldConfigArgument{
				"url": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"size": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"category": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				// marshall and cast the argument value
				url, _ := params.Args["url"].(string)
				size, _ := params.Args["size"].(int)
				category, _ := params.Args["category"].(string)

				db := initDb()
				defer db.Close()
				newImage := insertSrcImage(db, url, size, category)
				return newImage, nil
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
