package main

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"net/http"
	"testing"
	"github.com/unrolled/render"
	"bytes"
)

var (
	formatter = render.New(render.Options{
		IndentJSON: true})
)

func TestCreateMatch(t *testing.T) {
	client := &http.Client{}
	server := httptest.NewServer(http.HandlerFunc(createMatchHandler(formatter)))
	defer server.Close()
	body := []byte("{'test':123}")
	req, err := http.NewRequest("POST", server.URL, bytes.NewBuffer(body))
	if err != nil {
		t.Errorf("Error in create POST: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		t.Errorf("Error in POST: %v", err)
	}

	defer res.Body.Close()

	payload, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error in read POST: %v", err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("201 err: %v", err)
	}

	if _, ok := res.Header["Location"]; !ok {
		t.Error("Header Location not set")
	}

	fmt.Printf("payload: %s", string(payload))
}