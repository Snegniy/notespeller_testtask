package handlers

import (
	"context"
	"encoding/json"
	"github.com/Snegniy/notespeller-testtask/internal/model"
	"github.com/Snegniy/notespeller-testtask/pkg/logger"
	"go.uber.org/zap"
	"net/http"
)

type Handlers struct {
	srv Services
}

type Services interface {
	Login(ctx context.Context, username, password string) (string, error)
	Register(ctx context.Context, username, password string) error
	AddNote(ctx context.Context, note model.Note) (model.Note, error)
	GetNotes(ctx context.Context, userId int) ([]model.Note, error)
}

func NewHandlers(srv Services) *Handlers {
	logger.Debug("Handler:Creating handler")
	return &Handlers{srv: srv}
}

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) Register(w http.ResponseWriter, r *http.Request) {
	body := model.User{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		logger.Error("json Decoder failed", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.srv.Register(r.Context(), body.Username, body.Password)
	if err != nil {
		logger.Warn("Couldn't register user", zap.Error(err))
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	//writeJSON(w, res)
}

func (h *Handlers) AddNote(w http.ResponseWriter, r *http.Request) {
	body := model.Note{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		logger.Error("json Decoder failed", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := h.srv.AddNote(r.Context(), body)
	if err != nil {
		logger.Warn("Couldn't add note", zap.Error(err))
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	writeJSON(w, res)
}

func (h *Handlers) GetNotes(w http.ResponseWriter, r *http.Request) {
	res, err := h.srv.GetNotes(r.Context(), 1)
	if err != nil {
		logger.Warn("Couldn't get notes", zap.Error(err))
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	writeJSON(w, res)
}
