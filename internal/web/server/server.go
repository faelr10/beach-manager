package server

import (
	"go/beach-manager/internal/service"
	"go/beach-manager/internal/web/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	router *chi.Mux
	server *http.Server
	userService *service.UserService
	port string
}

func NewServer(userService *service.UserService, port string) *Server {

	return &Server{
		router: chi.NewRouter(),
		userService: userService,
		port: port,
	}
}

func (s *Server) ConfigureRoutes() {
	userHandler := handlers.NewUserHandler(s.userService)
	s.router.Post("/users", userHandler.CreateUser)
}

func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}

	s.ConfigureRoutes()

	return s.server.ListenAndServe()
}