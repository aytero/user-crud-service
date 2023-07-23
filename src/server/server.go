package server

import (
    "context"
    "net/http"
    "time"
    "user-crud-service/config"
)

const (
    _defaultAddr            = ":8080"
    _defaultReadTimeout     = 5 * time.Second
    _defaultWriteTimeout    = 5 * time.Second
    _defaultShutdownTimeout = 3 * time.Second
)

type Option func(*Server)

type Server struct {
    cfg             config.HTTP
    server          *http.Server
    notify          chan error
    shutdownTimeout time.Duration
}

func NewServer(cfg config.HTTP, handler http.Handler, opts ...Option) *Server {
    httpServer := &http.Server{
        Handler:      handler,
        ReadTimeout:  _defaultReadTimeout,
        WriteTimeout: _defaultWriteTimeout,
        Addr:         ":" + cfg.Port,
    }
    s := &Server{
        server:          httpServer,
        notify:          make(chan error, 1),
        shutdownTimeout: _defaultShutdownTimeout,
    }

    for _, opt := range opts {
        opt(s)
    }

    //s.SetupMiddleware(s.server)

    return s
}

func (s *Server) Start() {
    go func() {
        s.notify <- s.server.ListenAndServe()
        close(s.notify)
    }()
}

// Notify returns server's error chan
func (s *Server) Notify() <-chan error {
    return s.notify
}

// Shutdown sets up timer for shutdown and sends a signal through context.Context
func (s *Server) Shutdown(ctx context.Context) error {
    ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
    defer cancel()
    return s.server.Shutdown(ctx)
}
