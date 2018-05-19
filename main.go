package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/mnmtanish/go-graphiql"
	"github.com/rs/cors"
)

type Person struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

var PersonList []Person

func init() {
	person1 := Person{ID: "a", Name: "Jenny"}
	person2 := Person{ID: "b", Name: "Henery"}
	person3 := Person{ID: "b", Name: "Abagle"}
	PersonList = append(PersonList, person1, person2, person3)

	rand.Seed(time.Now().UnixNano())
}

// var People []Person
func getPeople() []Person {
	results := PersonList
	return results
}

var personType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Person",
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
		"People": &graphql.Field{
			Type:        graphql.NewList(personType),
			Description: "List of people",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return getPeople(), nil
			},
		},
	},
})

// root mutation
var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		/*
			curl -g 'http://localhost:8080/graphql?query=mutation+_{createPerson(text:"My+new+person"){id,text,done}}'
		*/
		"createPerson": &graphql.Field{
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
				// for e.g. create a Person and save to DB.
				newPerson := Person{
					ID:   RandStringRunes(8),
					Name: name,
				}

				PersonList = append(PersonList, newPerson)

				// return the new Person object that we supposedly save to DB
				// Note here that
				// - we are returning a `Person` struct instance here
				// - we previously specified the return Type to be `personType`
				// - `Person` struct maps to `personType`, as defined in `personType` ObjectConfig`
				return newPerson, nil
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
// 	People {
// 	  name
// 	  id
// 	}
//   }

// mutation {
// 	createPerson(name: "Lenny") {
// 	  name
// 	  id
// 	}
//   }
