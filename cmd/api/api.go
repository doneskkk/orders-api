package api

import (
	"database/sql"
	"github.com/doneskkk/order-api/service/user"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type apiServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *apiServer {
	return &apiServer{
		addr: addr,
		db:   db,
	}
}

func (s *apiServer) Run() error {
	router := chi.NewRouter()
	subrouter := chi.NewRouter()
	router.Mount("/api/v1", subrouter)

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)
	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
