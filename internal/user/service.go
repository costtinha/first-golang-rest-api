package user

import (
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Service interface {
	Create(ctx context.Context, in CreateUserInput) (*User, error)
	GetById(ctx context.Context, id uuid.UUID) (*User, error)
	List(ctx context.Context, page, size int) ([]User, int64, error)
	Update(ctx context.Context, id uuid.UUID, in UpdateUserInput) (*User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type service struct {
	repo      Repository
	validator *validator.Validate
}

func NewService(repo Repository) Service {
	return &service{repo: repo, validator: validator.New()}
}

func (s *service) Create(ctx context.Context, in CreateUserInput) (*User, error) {
	if err := s.validator.Struct(in); err != nil {
		return nil, err
	}
	u := &User{Name: in.Name, Email: in.Email}
	if err := s.repo.Create(ctx, u); err != nil {
		return nil, err
	}
	return u, nil

}

func (s *service) GetById(ctx context.Context, id uuid.UUID) (*User, error) {
	return s.repo.FindById(ctx, id)
}

func (s *service) List(ctx context.Context, page, size int) ([]User, int64, error) {
	if size <= 10 {
		size = 10

	}
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * size
	return s.repo.List(ctx, size, offset)
}

func (s *service) Update(ctx context.Context, id uuid.UUID, in UpdateUserInput) (*User, error) {
	if err := s.validator.Struct(in); err != nil {
		return nil, err
	}
	u, err := s.repo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	if in.Name != nil {
		u.Name = *in.Name
	}
	if in.Email != nil {
		u.Email = *in.Email
	}
	if err := s.repo.Update(ctx, u); err != nil {
		return nil, err
	}
	return u, nil
}

func (s *service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

var ErrNoFound = errors.New("user not found")
