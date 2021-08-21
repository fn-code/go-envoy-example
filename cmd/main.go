package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"github.com/fn-code/go-envoy-example/handler"
)

func main() {

	args := os.Args
	if len(args) < 2 {
		log.Fatal("invalid arguments")
	}

	p, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	h := &handler.Server{
		Port: p,
	}

	mux.Handle("/", h)

	log.Printf("server running on port:%d", p)
	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%d", p),
		Handler: mux,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
