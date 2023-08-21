package postgres

import (
	"context"
	"errors"
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

	time.Sleep(time.Second * 3)

	db, err := sqlx.Connect("pgx", cfg.Postgres.ConnString)
	if err != nil {
		logger.Fatal("not connected to base", zap.Error(err), zap.String("connection", cfg.Postgres.ConnString))
		return nil, err
	}
	return &Storage{db: db}, nil
}

func (r *Storage) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	logger.Debug("Repo:Login to PostgresSQL storage", zap.String("username", username))

	var res model.User
	sql := "SELECT * FROM users WHERE username = $1"

	if err := r.db.GetContext(ctx, &res, sql, username); err != nil {
		logger.Warn("login failed", zap.String("username", username))
		return model.User{}, errors.New("user not found")
	}
	return res, nil
}

func (r *Storage) AddUser(ctx context.Context, username, password string) error {
	logger.Debug("Repo:Register to PostgresSQL storage", zap.String("username", username), zap.String("password_hash", password))

	sql := "INSERT INTO users (username,password_hash) values ($1,$2)"

	if _, err := r.db.ExecContext(ctx, sql, username, password); err != nil {
		logger.Warn("Couldn't register user", zap.Error(err))
		return errors.New("user already exists")
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
		return nil, err
	}
	return res, nil
}
