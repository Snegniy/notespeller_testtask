package service

import (
	"context"
	"github.com/Snegniy/notespeller-testtask/internal/model"
	"github.com/Snegniy/notespeller-testtask/pkg/logger"
	"go.uber.org/zap"
)

type Service struct {
	repo Storage
}

type Storage interface {
	Login(ctx context.Context, username, password string) (string, error)
	Register(ctx context.Context, username, password string) error
	AddNote(ctx context.Context, note model.Note) (model.Note, error)
	GetNotes(ctx context.Context, userId int) ([]model.Note, error)
}

func NewService(repo Storage) *Service {
	logger.Debug("Service:Creating service")
	return &Service{repo: repo}
}

func (s *Service) Login(ctx context.Context, username, password string) (string, error) {
	return s.repo.Login(ctx, username, password)
}

func (s *Service) Register(ctx context.Context, username, password string) error {
	hashedPassword, err := generatePasswordHash(password)
	if err != nil {
		logger.Warn("Couldn't generate password hash", zap.Error(err))
		return err
	}
	return s.repo.Register(ctx, username, hashedPassword)
}

func (s *Service) AddNote(ctx context.Context, note model.Note) (model.Note, error) {
	err := speller(note.Note)
	if err != nil {
		logger.Warn("Couldn't check note", zap.Error(err))
		return model.Note{}, err
	}

	res, err := s.repo.AddNote(ctx, note)
	if err != nil {
		logger.Warn("Couldn't add note", zap.Error(err))
		return model.Note{}, err
	}
	return res, nil
}

func (s *Service) GetNotes(ctx context.Context, userId int) ([]model.Note, error) {
	return s.repo.GetNotes(ctx, userId)
}
