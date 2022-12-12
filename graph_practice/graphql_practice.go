package graph_practice

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type Tutorial struct {
	ID       int
	Title    string
	Author   Author
	Comments []Comment
}

type Author struct {
	Name      string
	Tutorials []int
}

type Comment struct {
	Body string
}

func GetGraphqlHandler() *handler.Handler {
	tutorials := populate()

	commentType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Comment",
		Fields: graphql.Fields{
			"body": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
	authorType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Author",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"tutorials": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	})
	tutorialType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Tutorial",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"author": &graphql.Field{
				Type: authorType,
			},
			"comments": &graphql.Field{
				Type: graphql.NewList(commentType),
			},
		},
	})

	fields := graphql.Fields{
		"tutorial": &graphql.Field{
			Type:        tutorialType,
			Description: "Get tutorials by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["id"].(int)
				if ok {
					for _, tutorial := range tutorials {
						if int(tutorial.ID) == id {
							return tutorial, nil
						}
					}
					return nil, errors.New("not tutorial found")
				}
				return "World", nil
			},
		},
		"list": &graphql.Field{
			Type:        graphql.NewList(tutorialType),
			Description: "Get the tutorial list",
			Resolve: func(q graphql.ResolveParams) (interface{}, error) {
				return tutorials, nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new graphql schema, err: %v", err.Error())
	}

	if false == true {
		runManualQuery(&schema)
	}

	gqlHandler := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})
	return gqlHandler
}

func populate() []Tutorial {
	author := &Author{Name: "Elliot Forbes", Tutorials: []int{1}}
	tutorial := Tutorial{
		ID:       1,
		Title:    "Go GraphQL tutorial",
		Author:   *author,
		Comments: []Comment{{Body: "First comment"}},
	}
	var tutorials []Tutorial
	tutorials = append(tutorials, tutorial)

	return tutorials
}

func runManualQuery(schema *graphql.Schema) {
	query := `
	{
		tutorial(id:1) {
			id
			title
			author {
				name
				tutorials
			}
		}
	}`

	params := graphql.Params{Schema: *schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute the graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)
}
