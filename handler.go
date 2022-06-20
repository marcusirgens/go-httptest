package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Bar struct {
	Baz string
}

type MyExternalService interface {
	Foo(string) (Bar, error)
}

type Handler struct {
	Ext MyExternalService
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cts, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	res, err := h.Ext.Foo(string(cts))
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(fmt.Errorf("formatting response: %w", err))
	}
}
