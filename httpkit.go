package main

import (
	"fmt"
	"net/http"
	"log"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		log.Print(r.URL.Path)
		next.ServeHTTP(wr, r)
	})
}

func VerifyHost(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("user") == "otto" {
			next.ServeHTTP(wr, r)
		} else {
			wr.Write([]byte("500 error"))
		}
	})
}

type middleware func(http.Handler) http.Handler

type Pipeline struct {
	middlewares []middleware
	handler http.Handler
}

func (p *Pipeline) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	h := p.handler

	for _, m := range p.middlewares {
		h = m(h)
	}

	h.ServeHTTP(wr, r)
}

func (p *Pipeline) Handler(handler http.Handler) *Pipeline {
	p.handler = handler;

	return p
}

func (p *Pipeline) With(pipeline *Pipeline) *Pipeline {
	return &Pipeline{
		middlewares: append(pipeline.middlewares...),
	}
}

func NewPipeline(middlewares ...middleware) *Pipeline {
	return &Pipeline{
		middlewares: middlewares,
	}
}


func main() {
	fmt.Println("starting server on localhost:3000")

	foo := http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		wr.Write([]byte("hell yea"))
	})

	pipeline := NewPipeline(Logger, VerifyHost)

	pipeline.Handler(foo)

	http.Handle("/foo", pipeline)
	
	http.ListenAndServe("localhost:3000", nil)
}

