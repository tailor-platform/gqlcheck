package main

import (
	"testing"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/tailor-platform/gqlcheck"
	"github.com/tailor-platform/gqlgen-todos/graph"
)

func TestServer(t *testing.T) {
	h := handler.New(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.Resolver{},
			},
		),
	)
	h.AddTransport(transport.POST{})

	checker := gqlcheck.New(h, gqlcheck.Debug())
	checker.Test(t).
		WithHeader("Content-Type", "application/json").
		Query(`query {todos {text}}`).
		Check().
		HasStatusOK().
		HasNoErrors().
		HasData(map[string]any{
			"todos": []any{},
		})
}
