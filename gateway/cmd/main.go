package main

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/tuanngoo192003/gateway-demo-go/gateway/internal/proxy"
	"github.com/tuanngoo192003/gateway-demo-go/gateway/internal/server"
)

func main() {
	s := server.NewServer()
	r := server.GetRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Mount("/", proxy.AuthAPIProxy())
	s.ListenAndServe()
}
