package main

import (
	"net/http"
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

func TestServerViaGet(t *testing.T) {
	h := handler.New(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.Resolver{},
			},
		),
	)
	h.AddTransport(transport.GET{})

	checker := gqlcheck.New(h, gqlcheck.Debug())
	checker.TestWithMethod(t, http.MethodGet).
		Query(`query {todos {text}}`).
		Check().
		HasStatusOK().
		HasNoErrors().
		HasData(map[string]any{
			"todos": []any{},
		})
}

func TestServerViaGetWithVariables(t *testing.T) {
	h := handler.New(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.Resolver{},
			},
		),
	)
	h.AddTransport(transport.GET{})
	h.AddTransport(transport.POST{})

	checker := gqlcheck.New(h, gqlcheck.Debug())

	// Create a todo via POST mutation with variables
	checker.Test(t).
		QueryWithVariables(
			`mutation CreateTodo($input: NewTodo!) { createTodo(input: $input) { id text } }`,
			map[string]any{"input": map[string]any{"text": "test todo", "userId": "user1"}},
		).
		Check().
		HasStatusOK().
		HasNoErrors()

	// Query via GET (variables passed to demonstrate the feature)
	checker.TestWithMethod(t, http.MethodGet).
		QueryWithVariables(
			`query { todos { text } }`,
			map[string]any{},
		).
		Check().
		HasStatusOK().
		HasNoErrors()
}
