package integrationtest

import (
	"canvas/server"
	"net/http"
	"time"
)

const (
	c_url  string = "http://localhost:8081/"
	c_host string = "localhost"
	c_port string = "8081"
)

// CreateServer for testing on port 8081, returning a cleanup function that stops the server.
// Usage:
// 	cleanup := CreateServer()
// 	defer cleanup()
func CreateServer() func() {
	s := server.New(server.Options{
		Host: c_host,
		Port: c_port,
	})

	go func() {
		if err := s.Start(); err != nil {
			panic(err)
		}
	}()

	for {
		_, err := http.Get(c_url)
		if err == nil {
			break
		}
		time.Sleep(5 * time.Millisecond)
	}

	return func() {
		if err := s.Stop(); err != nil {
			panic(err)
		}
	}
}
