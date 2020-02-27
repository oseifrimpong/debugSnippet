package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloHandler(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/", nil)

	helloHandler(res, req)

	content, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, "Hello, World!", string(content))
}
