package task

import (
	"context"
	"errors"

	"taskflow/internal/project"
)

type Service struct {
	repo        *Repository
	projectRepo *project.Repository
}

func NewService(repo *Repository, projectRepo *project.Repository) *Service {
	return &Service{repo: repo, projectRepo: projectRepo}
}

func (s *Service) Create(ctx context.Context, t *Task, userID string) error {
	p, err := s.projectRepo.GetByID(ctx, t.ProjectID)
	if err != nil {
		return err
	}

	if p.OwnerID != userID {
		return errors.New("forbidden")
	}

	if t.Status == "" {
		t.Status = "todo"
	}

	if t.Priority == "" {
		t.Priority = "medium"
	}

	return s.repo.Create(ctx, t)
}

func (s *Service) GetByProject(ctx context.Context, projectID, status, assignee, userID string) ([]Task, error) {
	p, err := s.projectRepo.GetByID(ctx, projectID)
	if err != nil {
		return nil, err
	}

	if p.OwnerID != userID {
		return nil, errors.New("forbidden")
	}

	return s.repo.GetByProject(ctx, projectID, status, assignee)
}

func (s *Service) Update(ctx context.Context, t *Task, userID string) error {
	existing, err := s.repo.GetByID(ctx, t.ID)
	if err != nil {
		return err
	}

	if t.Title == "" {
		t.Title = existing.Title
	}

	if t.Description == nil {
		t.Description = existing.Description
	}

	if t.Status == "" {
		t.Status = existing.Status
	}

	if t.Priority == "" {
		t.Priority = existing.Priority
	}

	if t.AssigneeID == nil {
		t.AssigneeID = existing.AssigneeID
	}

	if t.DueDate == nil {
		t.DueDate = existing.DueDate
	}

	p, err := s.projectRepo.GetByID(ctx, existing.ProjectID)
	if err != nil {
		return err
	}

	if p.OwnerID != userID {
		return errors.New("forbidden")
	}

	return s.repo.Update(ctx, t)
}

func (s *Service) Delete(ctx context.Context, id, userID string) error {
	t, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	p, err := s.projectRepo.GetByID(ctx, t.ProjectID)
	if err != nil {
		return err
	}

	if p.OwnerID != userID {
		return errors.New("forbidden")
	}

	return s.repo.Delete(ctx, id)
}
