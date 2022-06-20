package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type BarClient struct {
	Endpoint string
}

func (b BarClient) Foo(s string) (Bar, error) {
	payload := struct {
		ID string `json:"baz_id"`
	}{
		ID: s,
	}
	var buf bytes.Buffer

	// error handling intentionally omitted
	_ = json.NewEncoder(&buf).Encode(payload)
	req, _ := http.NewRequestWithContext(context.TODO(), http.MethodPost, b.Endpoint+"/v1/foo", &buf)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Bar{}, err
	}
	if res.StatusCode > 299 || res.StatusCode < 200 {
		return Bar{}, fmt.Errorf("got status code %d", res.StatusCode)
	}
	defer res.Body.Close()

	var response struct {
		Baz string `json:"baz"`
	}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return Bar{}, fmt.Errorf("parsing response: %w", err)
	}

	return Bar{Baz: response.Baz}, nil
}
