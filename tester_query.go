package gqlcheck

import (
	"encoding/json"
)

// Query is a struct to represent a query.
type Query struct {
	Query     string         `json:"query"`
	Variables map[string]any `json:"variables,omitempty"`
}

// String returns the string representation of the query.
func (q Query) String() string {
	b, _ := json.MarshalIndent(q, "", "  ") //nolint:errchkjson
	return string(b)
}

// Request sets the query and variables to the request.
func (tt *Tester) Request(q Query) *Tester {
	tt.query = q.Query
	tt.variables = q.Variables
	return tt
}

// Query sets the query to the request.
func (tt *Tester) Query(q string) *Tester {
	tt.query = q
	return tt
}

// QueryWithVariables sets the query and variables to the request.
func (tt *Tester) QueryWithVariables(q string, variables map[string]any) *Tester {
	tt.query = q
	tt.variables = variables
	return tt
}
