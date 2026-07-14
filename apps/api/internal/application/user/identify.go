package user

import (
	"context"
	"errors"

	"github.com/marcusantonius88/braqui/apps/api/internal/domain"
)

type UserRepository interface {
	FindByTelegramID(ctx context.Context, telegramID int64) (*domain.User, error)
	Create(ctx context.Context, user *domain.User) error
}

type IdentifyOutput struct {
	User    *domain.User
	IsNew   bool
}

type Identificer struct {
	repo UserRepository
}

func NewIdentificer(repo UserRepository) *Identificer {
	return &Identificer{repo: repo}
}

func (id *Identificer) Identify(ctx context.Context, telegramID int64, firstName, username string) (string, bool, error) {
	out, err := Identify(ctx, id.repo, telegramID, firstName, username)
	if err != nil {
		return "", false, err
	}
	return out.User.ID, out.IsNew, nil
}

func Identify(ctx context.Context, repo UserRepository, telegramID int64, firstName, username string) (*IdentifyOutput, error) {
	user, err := repo.FindByTelegramID(ctx, telegramID)
	if err == nil {
		return &IdentifyOutput{User: user, IsNew: false}, nil
	}

	if !errors.Is(err, domain.ErrNotFound) {
		return nil, err
	}

	user = &domain.User{
		TelegramID: telegramID,
		FirstName:  firstName,
		Username:   username,
	}

	if err := repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return &IdentifyOutput{User: user, IsNew: true}, nil
}
