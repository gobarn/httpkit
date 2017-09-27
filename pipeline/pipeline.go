package pipeline

import (
	"net/http"
)

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
	p.handler = handler

	return p
}

func (p *Pipeline) HandlerFunc(handler http.HandlerFunc) *Pipeline {
	p.handler = http.HandlerFunc(handler)

	return p
}

func (p *Pipeline) With(pipeline *Pipeline) *Pipeline {
	return &Pipeline{
		middlewares: pipeline.middlewares,
	}
}

func New(middlewares ...middleware) *Pipeline {
	return &Pipeline{
		middlewares: middlewares,
	}
}
