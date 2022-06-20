//go:build e2e

package main

import (
	"flag"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	extURL = flag.String("ext-url", "", "URL for external system")
)

func TestHandler_ServeHTTP_live(t *testing.T) {
	if *extURL == "" {
		t.Fatal("-ext-url is required")
	}

	h := Handler{Ext: BarClient{Endpoint: *extURL}}

	rec, req := httptest.NewRecorder(), httptest.NewRequest("POST", "/", strings.NewReader("hello, world!"))
	h.ServeHTTP(rec, req)
	res := rec.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("res.StatusCode = %d; want %d", res.StatusCode, http.StatusOK)
	}
}
