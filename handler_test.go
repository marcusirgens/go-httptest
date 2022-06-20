package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type extMock func(string) (Bar, error)

func (e extMock) Foo(s string) (Bar, error) {
	return e(s)
}

func TestHandler_ServeHTTP(t *testing.T) {
	mock := extMock(func(s string) (Bar, error) {
		return Bar{}, errors.New("remote service failure")
	})

	h := Handler{Ext: mock}

	rec, req := httptest.NewRecorder(), httptest.NewRequest("POST", "/", strings.NewReader("hello, world!"))
	h.ServeHTTP(rec, req)
	res := rec.Result()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("res.StatusCode = %d; want %d", res.StatusCode, http.StatusBadRequest)
	}
}
