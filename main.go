package main

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/mnmtanish/go-graphiql"
	"github.com/rs/cors"
)

type Person struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	From string `json:"from"`
}

// var People []Person
func getPeople() []Person {
	results := []Person{Person{1, "Jenny", "false"}, Person{2, "Henry", "true"}}
	return results
}

var personType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Person",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"from": &graphql.Field{
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
				// People = getPeople()
				return getPeople(), nil
			},
		},
	},
})
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queryType,
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
