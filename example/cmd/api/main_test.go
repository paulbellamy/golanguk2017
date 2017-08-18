package main

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"net"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	s := server(":0", ".")
	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		t.Fatal(err)
	}
	go s.Serve(ln)
	defer s.Shutdown(context.Background())

	// It should serve static assets
	if resp, err := http.Get(fmt.Sprintf("http://%s/", ln.Addr())); err != nil {
		t.Error(err)
	} else if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got: %s", resp.Status)
	}

	// It should 404 on unknown
	if resp, err := http.Get(fmt.Sprintf("http://%s/unknown", ln.Addr())); err != nil {
		t.Error(err)
	} else if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status code 404, got: %s", resp.Status)
	}

	// It should handle posts to /order
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "gopher.png")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := part.Write([]byte("hello world")); err != nil {
		t.Fatal(err)
	}
	for key, val := range map[string]string{
		"cut": "sirloin",
		"age": "21",
	} {
		if err := writer.WriteField(key, val); err != nil {
			t.Fatal(err)
		}
	}
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}
	if resp, err := http.Post(fmt.Sprintf("http://%s/order", ln.Addr()), writer.FormDataContentType(), body); err != nil {
		t.Error(err)
	} else if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got: %s", resp.Status)
	}
}
