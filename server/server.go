package server

import (
	"net/http"
	"server/globals"
	"server/server/router"
)

type Server struct {
	http.Server
}

func NewServer() *Server {
	return &Server{
		Server: http.Server{
			Addr:    ":" + globals.SERVER_PORT,
			Handler: router.NewRouter(),
		},
	}
}
