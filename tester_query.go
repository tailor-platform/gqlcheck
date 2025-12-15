package gqlcheck

import (
	"encoding/json"
	"net/http"
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

// Request sets the query and variables to the request (POST).
func (tt *Tester) Request(q Query) *Tester {
	tt.method = http.MethodPost
	tt.query = q.Query
	tt.variables = q.Variables
	return tt
}

// Query sets the query to the request (POST).
func (tt *Tester) Query(q string) *Tester {
	tt.method = http.MethodPost
	tt.query = q
	return tt
}

// QueryWithVariables sets the query and variables to the request (POST).
func (tt *Tester) QueryWithVariables(q string, variables map[string]any) *Tester {
	tt.method = http.MethodPost
	tt.query = q
	tt.variables = variables
	return tt
}

// RequestViaGet sets the query and variables to the request (GET).
func (tt *Tester) RequestViaGet(q Query) *Tester {
	tt.method = http.MethodGet
	tt.query = q.Query
	tt.variables = q.Variables
	return tt
}

// QueryViaGet sets the query to the request (GET).
func (tt *Tester) QueryViaGet(q string) *Tester {
	tt.method = http.MethodGet
	tt.query = q
	return tt
}

// QueryViaGetWithVariables sets the query and variables to the request (GET).
func (tt *Tester) QueryViaGetWithVariables(q string, variables map[string]any) *Tester {
	tt.method = http.MethodGet
	tt.query = q
	tt.variables = variables
	return tt
}
