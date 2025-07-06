package server

import (
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tuanngoo192003/gateway-demo-go/gateway/configs"
)

type Server struct {
	httpServer *http.Server
	router     *chi.Mux
}

var server *Server = nil

func NewServer() *Server {
	configs.Load()
	config := configs.GetConfig()

	if server == nil {
		r := chi.NewRouter()
		server = &Server{
			router: r,
			httpServer: &http.Server{
				Addr:    config.Server.Port,
				Handler: r,
			},
		}
	}
	return server
}

func (s *Server) ListenAndServe() {
	l, err := net.Listen("tcp", s.httpServer.Addr)
	if err != nil {
		panic(err)
	}
	s.httpServer.Serve(l)
}

func GetRouter() *chi.Mux {
	return server.router
}

func GetServer() *http.Server {
	return server.httpServer
}
