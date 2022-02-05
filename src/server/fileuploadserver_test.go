package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	//"gotest.tools/assert"
)

var request = struct {
	path string // request
}{
	path: "/list",
}

var (
	twitterResponse = `500`
)

func TestUploadServer(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		//io.WriteString(w, request.body)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	checkBody(t, resp, twitterResponse)
}

func checkBody(t *testing.T, r *http.Response, body string) {
	_, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Error("reading reponse body: , want given error")
	}
	if r.StatusCode != 200 {
		t.Errorf("for StatusCode = %d, want 200", r.StatusCode)
	}
}
