package main

import (
	"fmt"
	"io"
	"net/http"
)

type Route interface {
	http.Handler
	Pattern() string
}

type EchoHandler struct{}

type HelloHandler struct{}

func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

func (*HelloHandler) Pattern() string {
	return "/hello"
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	fmt.Fprintf(w, "Hello, %s\n", body)
}

func NewEchoHandler() *EchoHandler {
	return &EchoHandler{}
}

func (*EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = io.Copy(w, r.Body)
}

func (*EchoHandler) Pattern() string {
	return "/echo"
}

func NewServeMux(r1, r2 Route) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle(r1.Pattern(), r1)
	mux.Handle(r2.Pattern(), r2)
	return mux
}
