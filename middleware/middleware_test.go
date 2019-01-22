package middleware

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWithHeaders(t *testing.T) {
	ts := httptest.NewServer(WithHeaders(
		map[string]string{
			"Content-Type":      "application/json",
			"X-Ultimate-Header": "42",
		},
	)(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("{\"status\":\"ok\"}"))
			},
		),
	))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	assert.NoError(t, err)
	if res != nil {
		defer res.Body.Close()
	}
	assert.Equal(t, res.StatusCode, http.StatusOK)
	assert.Equal(t, res.Header.Get("Content-Type"), "application/json")
	assert.Equal(t, res.Header.Get("X-Ultimate-Header"), "42")
}
