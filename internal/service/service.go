package service

import (
	"context"
	"errors"
	"github.com/Snegniy/notespeller-testtask/internal/model"
	"github.com/Snegniy/notespeller-testtask/pkg/logger"
	"go.uber.org/zap"
)

type Service struct {
	repo Storage
}

type Storage interface {
	GetUserByUsername(ctx context.Context, username string) (model.User, error)
	AddUser(ctx context.Context, username, password string) error
	AddNote(ctx context.Context, note model.Note) (model.Note, error)
	GetNotes(ctx context.Context, userId int) ([]model.Note, error)
}

func NewService(repo Storage) *Service {
	logger.Debug("Service:Creating service")
	return &Service{repo: repo}
}

func (s *Service) UserLogin(ctx context.Context, username, password string) (model.User, error) {
	res, err := s.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return model.User{}, err
	}
	check := checkPasswordHash(res.Password, password)
	if !check {
		logger.Warn("login failed", zap.String("username", username))
		return model.User{}, errors.New("password not correct")
	}
	return res, nil
}

func (s *Service) UserRegister(ctx context.Context, username, password string) error {
	hashedPassword, err := generatePasswordHash(password)
	if err != nil {
		logger.Warn("Couldn't generate password hash", zap.Error(err))
		return err
	}
	return s.repo.AddUser(ctx, username, hashedPassword)
}

func (s *Service) AddNote(ctx context.Context, username string, content string) (model.Note, error) {
	err := speller(content)
	if err != nil {
		logger.Warn("Couldn't check note", zap.Error(err))
		return model.Note{}, err
	}

	userID, err := s.repo.GetUserByUsername(ctx, username)
	if err != nil {
		logger.Warn("Couldn't get user", zap.Error(err))
		return model.Note{}, err
	}

	note := model.Note{
		UserId: userID.Id,
		Note:   content,
	}

	res, err := s.repo.AddNote(ctx, note)
	if err != nil {
		logger.Warn("Couldn't add note", zap.Error(err))
		return model.Note{}, err
	}
	return res, nil
}

func (s *Service) GetNotes(ctx context.Context, username string) ([]model.Note, error) {
	userID, err := s.repo.GetUserByUsername(ctx, username)
	if err != nil {
		logger.Warn("Couldn't get user", zap.Error(err))
		return nil, err
	}
	return s.repo.GetNotes(ctx, userID.Id)
}
