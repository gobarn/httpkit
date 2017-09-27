package main

import (
	"fmt"
	"net/http"
	"github.com/gobarn/httpkit/pipeline"
	"github.com/gobarn/httpkit/auth"
	"github.com/gobarn/httpkit/logger"
)

func foo(wr http.ResponseWriter, r *http.Request) {
	wr.Header().Set("Content-Type", "text/html")
	wr.Write([]byte("<p>hell yeah ... this is our brand new homepage</p>"))
}

func main() {
	fmt.Println("starting server on localhost:3000")

	web := pipeline.New(
		logger.New,
		auth.New,
	)

	http.Handle("/foo", web.HandlerFunc(foo))
	
	http.ListenAndServe("localhost:3000", nil)
}

