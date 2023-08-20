package main

import (
	"github.com/Snegniy/notespeller-testtask/api"
	"github.com/Snegniy/notespeller-testtask/internal/config"
	"github.com/Snegniy/notespeller-testtask/internal/handlers"
	"github.com/Snegniy/notespeller-testtask/internal/service"
	"github.com/Snegniy/notespeller-testtask/internal/storage/postgres"
	"github.com/Snegniy/notespeller-testtask/pkg/graceful"
	"github.com/Snegniy/notespeller-testtask/pkg/logger"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
	"net/http"
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
	graceful.StartServer(r, cfg.ServerPort)

}
func Register(r *chi.Mux, h *handlers.Handlers) {
	r.Post("/login", h.Login)
	r.Post("/register", h.Register)
	r.Post("/", h.AddNote)
	r.Get("/", h.GetNotes)

	//SwaggerUI
	r.Get("/swagger", api.SwaggerUI)
	r.Get("/public/*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))).ServeHTTP(w, r)
	})
}
