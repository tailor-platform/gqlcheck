package gqlcheck

// WithHeader set header in the request.
func (tt *Tester) WithHeader(key, value string) *Tester {
	tt.headers[key] = value
	return tt
}

// WithHeaders sets header in the request.
func (tt *Tester) WithHeaders(headers map[string]string) *Tester {
	for k, v := range headers {
		tt.headers[k] = v
	}
	return tt
}
