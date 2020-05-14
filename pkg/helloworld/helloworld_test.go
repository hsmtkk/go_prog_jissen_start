package helloworld_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hsmtkk/go_prog_jissen_start/pkg/helloworld"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(helloworld.Handler))
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/test")
	assert.Nil(t, err, "should be nil")
	want := []byte("Hello World, test!")
	got, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err, "should be nil")
	resp.Body.Close()
	assert.Equal(t, want, got, "should be equal")
}
