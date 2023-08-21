package handlers

import (
	"context"
	"encoding/json"
	"github.com/Snegniy/notespeller-testtask/internal/model"
	"github.com/Snegniy/notespeller-testtask/pkg/logger"
	"github.com/go-chi/jwtauth/v5"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Handlers struct {
	srv       Services
	TokenAuth *jwtauth.JWTAuth
}

type InfoResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Services interface {
	UserLogin(ctx context.Context, username, password string) (model.User, error)
	UserRegister(ctx context.Context, username, password string) error
	AddNote(ctx context.Context, username string, content string) (model.Note, error)
	GetNotes(ctx context.Context, username string) ([]model.Note, error)
}

func NewHandlers(srv Services) *Handlers {
	logger.Debug("Handler:Creating handler")
	return &Handlers{srv: srv, TokenAuth: tokenAuth}
}

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	body := model.User{}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		logger.Error("json Decoder failed", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if body.Username == "" || body.Password == "" {
		logger.Warn("missing user or password")
		http.Error(w, "missing user or password", http.StatusBadRequest)
		return
	}

	res, err := h.srv.UserLogin(r.Context(), body.Username, body.Password)
	if err != nil {
		logger.Warn("Couldn't login user", zap.Error(err))
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	token := MakeToken(res.Username)
	http.SetCookie(w, &http.Cookie{
		HttpOnly: true,
		Expires:  time.Now().Add(1 * time.Hour),
		SameSite: http.SameSiteLaxMode,
		Name:     "jwt",
		Value:    token,
	})

	resp := InfoResponse{
		Status:  "success",
		Message: "login success",
	}

	writeJSON(w, resp)
}

func (h *Handlers) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		HttpOnly: true,
		Expires:  time.Now(),
		SameSite: http.SameSiteLaxMode,
		Name:     "jwt",
		Value:    "",
	})

	resp := InfoResponse{
		Status:  "success",
		Message: "logout success",
	}

	writeJSON(w, resp)
}

func (h *Handlers) Register(w http.ResponseWriter, r *http.Request) {
	body := model.User{}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		logger.Error("json Decoder failed", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if body.Username == "" || body.Password == "" {
		logger.Warn("missing user or password")
		http.Error(w, "missing user or password", http.StatusBadRequest)
		return
	}

	err = h.srv.UserRegister(r.Context(), body.Username, body.Password)
	if err != nil {
		logger.Warn("Couldn't register user", zap.Error(err))
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	resp := InfoResponse{
		Status:  "success",
		Message: "register success",
	}

	writeJSON(w, resp)
}

func (h *Handlers) AddNote(w http.ResponseWriter, r *http.Request) {
	body := model.Note{}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		logger.Error("json Decoder failed", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, claims, _ := jwtauth.FromContext(r.Context())
	username := claims["username"].(string)

	res, err := h.srv.AddNote(r.Context(), username, body.Note)
	if err != nil {
		logger.Warn("Couldn't add note", zap.Error(err))
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	writeJSON(w, res)
}

func (h *Handlers) GetNotes(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	username := claims["username"].(string)

	res, err := h.srv.GetNotes(r.Context(), username)
	if err != nil {
		logger.Warn("Couldn't get notes", zap.Error(err))
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if len(res) == 0 {
		resp := InfoResponse{
			Status:  "error",
			Message: "notes not found",
		}
		writeJSON(w, resp)
		return
	}

	writeJSON(w, res)
}
