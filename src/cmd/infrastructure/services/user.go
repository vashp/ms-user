package services

import (
	"context"
	"errors"
	"github.com/vashp/ms-user/src/cmd/domain"
	"github.com/vashp/ms-user/src/cmd/domain/usecases"
)

type UserService struct {
	repo usecases.UserRepository
}

func NewUserService(repo usecases.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(ctx context.Context, user *domain.User) error {
	if user.Email == "" {
		return errors.New("email is required")
	}

	return s.repo.Create(ctx, user)
}

func (s *UserService) GetByID(ctx context.Context, id int) (*domain.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) Update(ctx context.Context, user *domain.User) error {
	if user.Email == "" {
		return errors.New("email is required")
	}

	return s.repo.Update(ctx, user)
}

func (s *UserService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
