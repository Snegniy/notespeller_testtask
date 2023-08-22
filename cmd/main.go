package main

import (
	"github.com/Snegniy/notespeller-testtask/internal/config"
	"github.com/Snegniy/notespeller-testtask/internal/handlers"
	"github.com/Snegniy/notespeller-testtask/internal/service"
	"github.com/Snegniy/notespeller-testtask/internal/storage/postgres"
	"github.com/Snegniy/notespeller-testtask/pkg/logger"
	"github.com/Snegniy/notespeller-testtask/pkg/server"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth/v5"
	"go.uber.org/zap"
)

func main() {
	cfg := config.NewConfig()
	logger.Init(cfg.DebugMode)

	logger.Debug("Create router...")
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	repo, err := postgres.NewRepository(cfg)
	if err != nil {
		logger.Fatal("database not open", zap.Error(err))
	}

	s := service.NewService(repo)
	h := handlers.NewHandlers(s)

	Register(r, h)
	server.StartServer(r, cfg.ServerPort)

}
func Register(r *chi.Mux, h *handlers.Handlers) {
	r.Post("/login", h.Login)
	r.Post("/register", h.Register)
	r.Get("/logout", h.Logout)

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(h.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/note", h.AddNote)
		r.Get("/note", h.GetNotes)
	})

}
