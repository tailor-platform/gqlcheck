//go:generate go run github.com/99designs/gqlgen generate
package graph

import "github.com/tailor-platform/gqlgen-todos/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos []*model.Todo
}
