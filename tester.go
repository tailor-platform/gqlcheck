package gqlcheck

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/ikawaha/httpcheck"
)

// TestingT is an interface wrapper around *testing.T.
type TestingT interface {
	Errorf(format string, args ...any)
	FailNow()
}

// Tester represents the GraphQL tester.
type Tester struct {
	// For building request (before Check)
	checker   *Checker
	t         TestingT
	method    string
	headers   map[string]string
	query     string
	variables map[string]any

	// For response assertions (after Check)
	client *httpcheck.Tester
}

// Test starts a new test with the given *testing.T.
// The default HTTP method is POST.
func (c *Checker) Test(t TestingT) *Tester {
	return &Tester{
		checker: c,
		t:       t,
		method:  http.MethodPost, // default
		headers: make(map[string]string),
	}
}

// TestWithMethod starts a new test with the given *testing.T and HTTP method.
func (c *Checker) TestWithMethod(t TestingT, method string) *Tester {
	return &Tester{
		checker: c,
		t:       t,
		method:  method,
		headers: make(map[string]string),
	}
}

// Check makes request to built request object.
// After request is made, it saves response object for future assertions.
func (tt *Tester) Check() *Tester {
	var client *httpcheck.Tester

	switch tt.method {
	case http.MethodGet:
		// GET: query parameters in URL
		params := url.Values{}
		params.Set("query", tt.query)
		if tt.variables != nil {
			v, err := json.Marshal(tt.variables)
			if err != nil {
				tt.t.Errorf("failed to marshal variables: %v", err)
				tt.t.FailNow()
			}
			params.Set("variables", string(v))
		}
		path := "?" + params.Encode()
		client = tt.checker.client.Test(tt.t, http.MethodGet, path)
	default:
		// POST: JSON body
		client = tt.checker.client.Test(tt.t, http.MethodPost, "").
			WithHeader("Content-Type", "application/json")
		body := map[string]any{"query": tt.query}
		if tt.variables != nil {
			body["variables"] = tt.variables
		}
		client = client.WithJSON(body)
	}

	// Apply headers
	for k, v := range tt.headers {
		client = client.WithHeader(k, v)
	}

	return &Tester{client: client.Check()}
}
