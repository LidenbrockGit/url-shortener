package defserver

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	srv http.Server
}

func NewServer(addr string, h http.Handler) *Server {
	s := &Server{}
	s.srv = http.Server{
		Addr:              addr,
		Handler:           h,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
	}
	return s
}

func (s *Server) Start() {
	go func() {
		err := s.srv.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	_ = s.srv.Shutdown(ctx)
	cancel()
}
