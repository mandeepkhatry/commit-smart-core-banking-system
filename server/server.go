package server

import (
	"fmt"
	"net/http"

	"github.com/commit-smart-core-banking-system/config"
	"github.com/commit-smart-core-banking-system/logger"
	"github.com/commit-smart-core-banking-system/routes"
)

type Server struct {
	Server *http.Server
}

func NewServer() *Server {
	srv := &http.Server{
		Addr:    fmt.Sprintf(config.AppConfiguration.ServerAddress),
		Handler: routes.RegisterRoutes(),
	}
	return &Server{
		Server: srv,
	}
}

func (s *Server) Run() error {
	logger.Debug("Running Server")
	return s.Server.ListenAndServe()
}
