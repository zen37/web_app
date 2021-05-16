package server_test

import (
	"net/http"
	"testing"

	"github.com/matryer/is"

	"canvas/integrationtest"
)

const (
	c_run string = "starts the server and listens for requests"
	c_url string = "http://localhost:8081/"
)

func TestServer_Start(t *testing.T) {
	t.Run(c_run, func(t *testing.T) {
		is := is.New(t)

		cleanup := integrationtest.CreateServer()
		defer cleanup()

		resp, err := http.Get(c_url)
		is.NoErr(err)
		is.Equal(http.StatusNotFound, resp.StatusCode)

	})
}
