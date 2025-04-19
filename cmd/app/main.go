package main

import (
	"database/sql"
	"fmt"
	"go/beach-manager/internal/provider"
	"go/beach-manager/internal/repository"
	"go/beach-manager/internal/service"
	"go/beach-manager/internal/web/server"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func getEnv(key, defaulValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaulValue
}

func main() {

	log.Println("Iniciando aplicação...")


	if os.Getenv("RENDER") == "" {
		// Só carrega o .env local se não estiver no Render
		if err := godotenv.Load(); err != nil {
			log.Println("Aviso: não foi possível carregar o arquivo .env (ambiente local)")
		}
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		getEnv("DB_HOST", "db"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "postgres"),
		getEnv("DB_SSLMODE", "require"),
	)

	log.Println("Conectando ao banco:", connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	secret := os.Getenv("JWT_SECRET")
	expMinutes, err := strconv.Atoi(os.Getenv("JWT_EXP_MINUTES"))
	if err != nil {
		log.Fatal("JWT_EXP_MINUTES is invalid")
	}

	jwtProvider := provider.NewJWTProvider(
		secret,
		time.Duration(expMinutes)*time.Minute,
		"minha_secret_refresh",
		7*24*time.Hour,
	)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	agendaRepository := repository.NewAgendaRepository(db)
	agendaService := service.NewAgendaService(agendaRepository)
	authService := service.NewAuthService(userRepository, *jwtProvider)

	port := getEnv("PORT", "8082")

	srv := server.NewServer(userService, agendaService, authService, jwtProvider, port)

	log.Println("Servidor iniciado na porta:", port)
	if err := srv.Start(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
