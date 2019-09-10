package blog

import (
	"github.com/anvh2/z-blogs/storages"
	"go.uber.org/zap"
)

// Server ...
type Server struct {
	blogDb storages.BlogDb
	logger zap.Logger
}

// NewServer ...
func NewServer(blogDb storages.BlogDb, logger zap.Logger) *Server {
	return &Server{
		blogDb: blogDb,
		logger: logger,
	}
}

// Run ...
func (s *Server) Run() error {
	return nil
}
