package server

import (
	"go/beach-manager/internal/service"
	"go/beach-manager/internal/web/handlers"
	"net/http"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	router        *chi.Mux
	server        *http.Server
	userService   *service.UserService
	agendaService *service.AgendaService
	authService   *service.AuthService
	port          string
}

func NewServer(userService *service.UserService, agendaService *service.AgendaService, authService *service.AuthService, port string) *Server {

	return &Server{
		router:        chi.NewRouter(),
		userService:   userService,
		agendaService: agendaService,
		authService:   authService,
		port:          port,
	}
}

func (s *Server) ConfigureRoutes() {

	userHandler := handlers.NewUserHandler(s.userService)
	agendaHandler := handlers.NewAgendaHandler(s.agendaService)
	authHandler := handlers.NewAuthHandler(s.authService)

	//Users
	s.router.Post("/users", userHandler.CreateUser)
	s.router.Get("/users/{id}", userHandler.GetById)

	//Agendas
	s.router.Post("/agendas", agendaHandler.CreateAgenda)
	s.router.Get("/agendas/{id}", agendaHandler.GetAgendaByID)
	s.router.Get("/agendas", agendaHandler.GetAllAgendas)

	//Auth
	s.router.Post("/auth/login", authHandler.Login)

}

func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}

	s.ConfigureRoutes()

	return s.server.ListenAndServe()
}
