package server

import (
	"context"
	"fmt"
	"net/http"
	"simplicity/llmbridge"
	"sync"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	ChiRouter chi.Router
	llmclient llmbridge.LLMClient
	sync.Mutex
}

func (s *Server) Start(port uint) {

	http.ListenAndServe(fmt.Sprintf(":%d", port), s.ChiRouter)

}

func NewServer(ctx context.Context, llmclient llmbridge.LLMClient) *Server {

	mux := chi.NewMux()
	return &Server{ChiRouter: mux, llmclient: llmclient}

}
