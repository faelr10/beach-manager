package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"go/beach-manager/internal/provider"
	"go/beach-manager/internal/service"
	"go/beach-manager/internal/web/handlers"
	"go/beach-manager/internal/web/middleware"
	"net/http"
)

type Server struct {
	router        *chi.Mux
	server        *http.Server
	userService   *service.UserService
	agendaService *service.AgendaService
	authService   *service.AuthService
	jwtProvider   *provider.JWTProvider
	port          string
}

func NewServer(userService *service.UserService, agendaService *service.AgendaService, authService *service.AuthService, jwtProvider *provider.JWTProvider, port string) *Server {

	return &Server{
		router:        chi.NewRouter(),
		userService:   userService,
		agendaService: agendaService,
		authService:   authService,
		jwtProvider:   jwtProvider,
		port:          port,
	}
}

func (s *Server) ConfigureRoutes() {
	// ✅ Adiciona CORS globalmente
	s.router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:3000", // para testes locais
			"https://front-beach-manager.onrender.com",
			"https://master.d3uhax1xjrnrnb.amplifyapp.com",
			"https://www.manager-sportings.com.br",
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	userHandler := handlers.NewUserHandler(s.userService)
	agendaHandler := handlers.NewAgendaHandler(s.agendaService)
	authHandler := handlers.NewAuthHandler(s.authService)

	// Grupo de rotas públicas
	s.router.Post("/users", userHandler.CreateUser)
	s.router.Get("/users/{id}", userHandler.GetById)

	s.router.Post("/auth/login", authHandler.Login)
	s.router.Post("/auth/refresh-token", authHandler.RefreshToken)

	//agenda-publica
	s.router.Get("/agenda-publica/{id}", agendaHandler.GetAllAgendasByUserIDFromURL)

	// Grupo de rotas protegidas
	s.router.Route("/agendas", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(s.jwtProvider)) // aplica o middleware
		r.Post("/", agendaHandler.CreateAgenda)
		// r.Get("/", agendaHandler.GetAllAgendas)
		r.Get("/", agendaHandler.GetAllAgendasByUserID)
		r.Get("/{id}", agendaHandler.GetAgendaByID)
		r.Put("/{id}", agendaHandler.UpdateAgenda)
		r.Delete("/{id}", agendaHandler.DeleteAgenda)
	})

	//rota teste
	s.router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

}

func (s *Server) Start() error {
	s.ConfigureRoutes()

	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}

	return s.server.ListenAndServe()
}
