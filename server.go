package todo

import (
	"context"
	"net/http"
	"time"
)

// для запуска Http-сервера
type Server struct {
	httpServer *http.Server
}

// methods
func (s *Server) Run(port string, handler http.Handler) error { // метод запуска
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe() // Под капотом бесконечный цикл For
}

func (s *Server) Shutdown(ctx context.Context) error { // при выходе из прлж
	return s.httpServer.Shutdown(ctx)
}
