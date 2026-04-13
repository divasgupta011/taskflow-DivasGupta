package project

import (
	"context"
	"errors"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, name string, desc *string, userID string) (*Project, error) {
	p := &Project{
		Name:        name,
		Description: desc,
		OwnerID:     userID,
	}

	err := s.repo.Create(ctx, p)
	return p, err
}

func (s *Service) GetAll(ctx context.Context, userID string) ([]Project, error) {
	return s.repo.GetAllByUser(ctx, userID)
}

func (s *Service) GetByID(ctx context.Context, id string, userID string) (*Project, error) {
	p, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if p.OwnerID != userID {
		return nil, errors.New("forbidden")
	}

	return p, nil
}

func (s *Service) Update(ctx context.Context, p *Project, userID string) error {
	existing, err := s.repo.GetByID(ctx, p.ID)
	if err != nil {
		return err
	}

	if existing.OwnerID != userID {
		return errors.New("forbidden")
	}

	return s.repo.Update(ctx, p)
}

func (s *Service) Delete(ctx context.Context, id string, userID string) error {
	p, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if p.OwnerID != userID {
		return errors.New("forbidden")
	}

	return s.repo.Delete(ctx, id)
}
