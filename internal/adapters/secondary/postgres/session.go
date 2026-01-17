package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/frostnzx/go-ecommerce-api/internal/core/domain/session"
	"github.com/jmoiron/sqlx"
)

type SessionRepo struct {
	db *sqlx.DB
}

func NewSessionRepo(db *sqlx.DB) (*SessionRepo, error) {
	if db == nil {
		return nil, errors.New("database connection required")
	}
	return &SessionRepo{db: db}, nil
}

func (sr *SessionRepo) CreateSession(ctx context.Context, s *session.Session) (*session.Session, error) {
	_, err := sr.db.NamedExecContext(ctx, "INSERT INTO sessions (id, user_email, refresh_token, is_revoked, expires_at) VALUES (:id, :user_email, :refresh_token, :is_revoked, :expires_at)", s)
	if err != nil {
		return nil, fmt.Errorf("error inserting session: %w", err)
	}
	return s, nil
}

func (sr *SessionRepo) GetSession(ctx context.Context, id string) (*session.Session, error) {
	var s session.Session
	err := sr.db.GetContext(ctx, &s, "SELECT * FROM sessions WHERE id=$1", id)
	if err != nil {
		return nil, fmt.Errorf("error getting session: %w", err)
	}
	return &s, nil
}

func (sr *SessionRepo) RevokeSession(ctx context.Context, id string) error {
	_, err := sr.db.ExecContext(ctx, "UPDATE sessions SET is_revoked=TRUE WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("error revoking session: %w", err)
	}
	return nil
}

func (sr *SessionRepo) DeleteSession(ctx context.Context, id string) error {
	_, err := sr.db.ExecContext(ctx, "DELETE FROM sessions WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("error deleting session: %w", err)
	}
	return nil
}
