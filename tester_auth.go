package gqlcheck

import "encoding/base64"

// WithBasicAuth is an alias to set basic auth in the request header.
func (tt *Tester) WithBasicAuth(user, pass string) *Tester {
	auth := base64.StdEncoding.EncodeToString([]byte(user + ":" + pass))
	tt.headers["Authorization"] = "Basic " + auth
	return tt
}

// WithBearerAuth is an alias to set bearer auth in the request header.
func (tt *Tester) WithBearerAuth(token string) *Tester {
	tt.headers["Authorization"] = "Bearer " + token
	return tt
}
