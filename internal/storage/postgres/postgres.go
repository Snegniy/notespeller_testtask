package postgres

import (
	"context"
	"github.com/Snegniy/notespeller-testtask/internal/apperrors"
	"github.com/Snegniy/notespeller-testtask/internal/config"
	"github.com/Snegniy/notespeller-testtask/internal/model"
	"github.com/Snegniy/notespeller-testtask/pkg/logger"
	_ "github.com/jackc/pgx"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"time"
)

type Storage struct {
	db *sqlx.DB
}

func NewRepository(cfg config.Config) (*Storage, error) {
	logger.Debug("Repo:Creating PostgresSQL repository")
	time.Sleep(time.Second * 2)
	db, err := sqlx.Connect("pgx", cfg.Postgres.ConnString)
	if err != nil {
		logger.Fatal("not connected to base", zap.Error(err), zap.String("connection", cfg.Postgres.ConnString))
		return nil, err
	}
	return &Storage{db: db}, nil
}

func (r *Storage) Login(ctx context.Context, username, password string) (string, error) {
	logger.Debug("Repo:Login to PostgresSQL storage", zap.String("username", username))
	var res string
	sql := "SELECT username FROM users WHERE username = $1 AND password_hash = $2"
	if err := r.db.GetContext(ctx, sql, username, password); err != nil {
		logger.Warn("login failed", zap.String("username", username))
		return "", apperrors.ErrUserNotFound
	}
	return res, nil
}

func (r *Storage) Register(ctx context.Context, username, password string) error {
	logger.Debug("Repo:Register to PostgresSQL storage", zap.String("username", username))
	sql := "INSERT INTO users (username,password_hash) values ($1,$2)"
	if _, err := r.db.ExecContext(ctx, sql, username, password); err != nil {
		logger.Warn("Couldn't register user", zap.Error(err))
		return err
	}
	return nil
}

func (r *Storage) AddNote(ctx context.Context, note model.Note) (model.Note, error) {
	logger.Debug("Repo:Add note to PostgresSQL storage", zap.String("content", note.Note))
	sql := "INSERT INTO notes (userid,note) values ($1, $2)"
	if _, err := r.db.ExecContext(ctx, sql, note.UserId, note.Note); err != nil {
		logger.Warn("Couldn't add note", zap.Error(err))
		return model.Note{}, err
	}

	res := model.Note{}
	sql = "SELECT * FROM notes WHERE userid = $1 ORDER BY id DESC LIMIT 1"
	if err := r.db.GetContext(ctx, &res, sql, note.UserId); err != nil {
		logger.Warn("Couldn't find note", zap.Error(err))
		return model.Note{}, err
	}
	return res, nil
}

func (r *Storage) GetNotes(ctx context.Context, userId int) ([]model.Note, error) {
	logger.Debug("Repo:Get notes from PostgresSQL storage", zap.Int("userId", userId))
	var res []model.Note
	sql := "SELECT * FROM notes WHERE userid = $1"
	if err := r.db.SelectContext(ctx, &res, sql, userId); err != nil {
		logger.Warn("Couldn't find note", zap.Error(err))
		return []model.Note{}, err
	}
	return res, nil
}
