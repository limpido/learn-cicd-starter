package auth

import (
	"testing"
	"net/http"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		name string
		input http.Header
		want string
		wantError error
	}
	
	headerEmpty := http.Header{
		"Host": {"boot.dev"},
	}
	headerMalformed := http.Header{
		"Authorization": {"Basic"},
	}
	headerCorrect := http.Header{
		"Authorization": {"ApiKey YWxhZGRpbjpvcGVuc2VzYW1l"},
	}

	tests := []test{
		{name: "empty header", input: headerEmpty, want: "", wantError: ErrNoAuthHeaderIncluded},
		{name: "malformed header", input: headerMalformed, want: "", wantError: ErrMalformedAuthHeader},
		{name: "correct header", input: headerCorrect, want: "YWxhZGRpbjpvcGVuc2VzYW1l", wantError: nil},
	}

	for _, tc := range tests {
		got, gotError := GetAPIKey(tc.input)
		if got != tc.want || gotError != tc.wantError {
			t.Fatalf("%s: expected: %v, %v, got: %v, %v", tc.name, tc.want, tc.wantError, got, gotError)
		}
	}
}
