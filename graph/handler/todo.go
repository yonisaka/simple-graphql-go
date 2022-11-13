package handler

import (
	"fmt"
	"simple-graphql-go/graph/schema"

	"github.com/graphql-go/graphql"
)

func TodoHandler(query string) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema.TodoSchema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
}
