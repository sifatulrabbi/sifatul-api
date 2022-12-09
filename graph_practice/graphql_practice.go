package graph_practice

import (
	"github.com/graphql-go/graphql"
	gqHandler "github.com/graphql-go/handler"
)

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"latestPost": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "Hello world!", nil
			},
		},
	},
})

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queryType,
})

var GqHandler = gqHandler.New(&gqHandler.Config{
	Schema: &schema,
	Pretty: true,
})
