package handler

import (
	"fmt"
	"net/http"
)

type Server struct {
	Port int
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello from port %d \n", s.Port)
}