package http_server

import (
	"backend-api-test/config"
	"context"
	"net/http"
	"time"
)

const (
	_defaultReadTimeout     = 5 * time.Second
	_defaultWriteTimeout    = 5 * time.Second
	_defaultAddr            = ":80"
	_defaultShutdownTimeout = 3 * time.Second
)

// Server -.
type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

// New -.
func Init(conf *config.Config, handler http.Handler, opts ...Option) *Server {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  conf.Server.ReadTimeout * time.Second,
		WriteTimeout: conf.Server.WriteTimeout * time.Second,
		Addr:         conf.Server.Port,
	}

	s := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: conf.Server.ShutdownTimeout * time.Second,
	}

	// Custom options
	for _, opt := range opts {
		opt(s)
	}

	s.start()

	return s
}

func (s *Server) start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

// Notify -.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown -.
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
