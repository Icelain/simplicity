package server

import (
	"context"
	"encoding/json"
	"net/http"
	//"simplicity/llmbridge"
	"simplicity/types/requestypes"

	"github.com/go-chi/chi/v5/middleware"
)

func ManageRoutes(ctx context.Context, srv *Server) {

	srv.ChiRouter.Use(
		middleware.Logger,
		middleware.RealIP,
		middleware.Recoverer,
	)

	srv.ChiRouter.Get("/", handleIndex)
	srv.ChiRouter.Get("/index", handleIndex)
	srv.ChiRouter.Post("/query", handleQuery(srv))

}

func handleIndex(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Todo"))

}

func handleQuery(srv *Server) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var q requestypes.Query

		defer r.Body.Close()

		if err := json.NewDecoder(r.Body).Decode(&q); err != nil {

			http.Error(w, "invalid json", http.StatusBadRequest)
			return

		}

		flusher, ok := w.(http.Flusher)
		if !ok {

			http.Error(w, "Could not create http stream", http.StatusInternalServerError)
			return

		}

		recvdata := make(chan string)

		srv.Lock()
		go srv.llmclient.StreamResponse(context.Background(), q.SearchQuery, recvdata)
		srv.Unlock()

		for {

			content, ok := <-recvdata

			if !ok {

				break

			}

			if _, err := w.Write([]byte(content)); err != nil {

				break

			}
			flusher.Flush()

		}

	})

}
